package src

import (
	"github.com/joernweissenborn/aursir4go"
	"github.com/joernweissenborn/aursir4go/messages"
	"log"
)

func InitServer(Storage *AurHobasStorage)  {

	log.Println("Server initializing")

	var server AurHobasServer
	server.iface, _ = aursir4go.NewInterface("AurHobasServer")
	server.storage = Storage

	server.initExporter()

	server.export = server.iface.AddExport(StoreAppkey,[]string{})

	log.Println("Server initialized")

	server.run()
}

type AurHobasServer struct {
	iface aursir4go.AurSirInterface
	export *aursir4go.ExportedAppKey
	storage *AurHobasStorage
	dataexports map[string] AurHobasDataExporter
}

func (ahs *AurHobasServer) initExporter(){
	ahs.dataexports = map[string]AurHobasDataExporter{}
	ahs.storage.IterateItems(ahs.AddExporter)

}

func (ahs *AurHobasServer) AddExporter(item AurHobasItem) {

	if _,f := ahs.dataexports[item.Id]; !f {

		ahs.dataexports[item.Id] = createDataExporter(ahs.iface, item)

	}

}


func (ahs AurHobasServer) run(){

	log.Println("Server running")


	for request := range ahs.export.Request {
		ahs.export.Reply(request,ahs.onRequest(request))
	}

}

func (ahs AurHobasServer) onRequest(request messages.Request) (reply interface {}){

	log.Println("Server Got Request",request.FunctionName)


	switch request.FunctionName{

	case "StoreItem":
		var sir StoreItemRequest
		request.Decode(&sir)
		log.Println("AddingData",sir.Name,sir.Tags,sir.Type)
		id := ahs.storage.AddItem(sir.Name,sir.Tags,sir.Type,sir.Data)
		reply = StoreItemReply{id}
		if ahs.storage.Exists(id){
			ahs.AddExporter(ahs.storage.Get(id))
		}
	}
	return
}
