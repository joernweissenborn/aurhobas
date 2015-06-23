package storage

import (
	"time"
	"fmt"
	"os"
	"log"
	"path"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"github.com/joernweissenborn/stream2go"
	"aurhobas/storage"
)

const META_FILE_NAME = "aurhobas.meta"


func CreatePathIfNecessary(path string){
	dir, err := os.Open(path)
	if err != nil {
		err =os.MkdirAll(path,os.ModeDir)
		handleDataStoreError(err)
	} else {
		dir.Close()
	}

}

func LoadData(path string, events stream2go.StreamController) {
	err := filepath.Walk(path, func(p string, f os.FileInfo, err error) error {
			handleDataStoreError(err)
			if f.Name() == META_FILE_NAME {
				events.Add(NewItem{p})
			}
			return nil
		})
	handleDataStoreError(err)
}

func handleDataStoreError(err error){
	if err !=nil {
		log.Fatal("DataStore ",err)
	}
}
