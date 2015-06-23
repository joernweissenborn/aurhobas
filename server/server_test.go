package src

import (
	"testing"
	"github.com/joernweissenborn/aursir4go"
)

var (
	testitem = StoreItemRequest{
		"testdata",
		"YAML",
		[]string{"Tag1","Tag2"},
		`dataname: bla
		array: [0,1,2,3]`,
		}

	ID = ""

	iface, _ = aursir4go.NewInterface("testaurhobas")
	imp = iface.AddImport(StoreAppkey, []string{})
	itemimp *aursir4go.ImportedAppKey
)



func TestStoreItem(T *testing.T){


	if !imp.Exported() {
		T.Fatal("AurHobas Offline")
	}

	req, err := imp.Call("StoreItem",testitem)

	if err != nil {
		T.Fatal(err)
	}

	var res StoreItemReply
	err = (<-req).Decode(&res)


	if err != nil {
		T.Fatal(err)
	}


	if res.Id == "" {
		T.Error("Did not get Id")
	}

	ID = res.Id
}

func TestConnectToItem(T *testing.T){

	itemimp = iface.AddImport(ItemAppkey,[]string{"testdata",ID,"YAML","Tag1","Tag2"})


	if !itemimp.Exported() {
		T.Fatal("Could not connect to data")
	}

}

func TestGetId(T *testing.T){

	r,_ := itemimp.Call("GetId",nil)
	var res GetIdReply
	(<-r).Decode(&res)

	if res.Id != ID {
		T.Errorf("Got Wrong Id, needed %s, got %s",ID, res.Id)
	}

}


func TestGetMeta(T *testing.T){

	r,_ := itemimp.Call("GetMeta",nil)
	var res GetMetaReply
	(<-r).Decode(&res)

	if res.Name != testitem.Name {
		T.Errorf("Got Wrong Name, needed %s, got %s",testitem.Name, res.Name)
	}

	if res.Type != testitem.Type {
		T.Errorf("Got wrong type, needed %s, got %s",testitem.Type, res.Type)
	}
	if !testEq(res.Tags,testitem.Tags) {
		T.Errorf("Got wrong tags, needed %s, got %s",testitem.Tags, res.Tags)
	}

}

func testEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
