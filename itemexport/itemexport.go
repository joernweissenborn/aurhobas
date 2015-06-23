package itemexport

import (
	"github.com/joernweissenborn/aursir4go"
	"github.com/joernweissenborn/aursir4go/messages"
	"io/ioutil"
	"log"
	"github.com/joernweissenborn/aursir4go/aurmellon"
	"aurhobas/item"
)

type ItemExporter struct {
	export aurmellon.AurMellonExport
	item item.Item
}

func New(i aurmellon.AurMellonInterface, metapath string) (ie ItemExporter){


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
