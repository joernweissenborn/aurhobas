package item
import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"aurhobas/storage"
)

type Item struct {
	Id string
	Name string
	Files []File
}

func NewFromFile(path string) (i Item, err error){

	var m Meta

	mf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("Error opening Meatafile at",path,err)
		return
	}

	err = yaml.Unmarshal(mf,m)
	if err != nil {
		log.Println("Error opening Meatafile at",path,err)
		return
	}

	i.Id = m.Id
	i.Name = i.Name
	folder, err := ioutil.ReadDir(filepath.Dir(path))
	if err != nil {
		log.Println("Error opening Meatafile at",path,err)
		return
	}
	for file := range folder {
		if !strings.Contains(file.Name(),storage.META_FILE_NAME) {

		}
	}

}