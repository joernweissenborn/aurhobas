package src

import (
	"github.com/joernweissenborn/aursir4go"
	"github.com/joernweissenborn/aursir4go/messages"
	"io/ioutil"
	"log"
)

type AurHobasDataExporter struct {
	export *aursir4go.ExportedAppKey
	meta AurHobasItem
	remove chan struct{}
}

func createDataExporter(i aursir4go.AurSirInterface, item AurHobasItem) AurHobasDataExporter{
	var e AurHobasDataExporter

	tags := []string{item.Id,item.Name,item.Type}

	for _, t := range item.Tags {
		tags = append(tags,t)
	}

	e.export = i.AddExport(ItemAppkey, tags)
	e.item = item
	go e.run()

	return e
}

func (ade AurHobasDataExporter) run() {
	for {
		select {

		case <-ade.remove:
			close(ade.remove)
			return

		case req, ok := <-ade.export.Request:
			if ok {
				ade.export.Reply(req, ade.request(req))
			}
		}
	}
}

func (ade AurHobasDataExporter) Remove(){
	ade.remove <- struct{}{}
}

func (ade AurHobasDataExporter) request(r messages.Request) (reply interface {}){

	switch r.FunctionName{

	case "GetId":
		reply = GetIdReply{ade.item.Id}

	case "GetMeta":
		reply = GetMetaReply{ade.item.Name,ade.item.Type,ade.item.Tags}

	case "GetData":
		reply = GetDataReply{ade.getData()}
	}
	return

}

func (ade AurHobasDataExporter) getData() (data string) {
	d, err := ioutil.ReadFile(ade.item.Datapath)
	data = string(d)
	if err != nil {
		log.Fatal(err)
	}
	return
}
