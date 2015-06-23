package src

import (
	"os"
	"log"
	"path/filepath"
	"io/ioutil"
	yaml "gopkg.in/yaml.v2"

)

type AurHobasConfig struct {
	RepositoryPath string
}

func GetConfig() (cfg AurHobasConfig) {

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	cfgpath := filepath.Join(cwd,"config.yaml")

	cfgfile, err := ioutil.ReadFile(cfgpath)
	if err != nil {
		cfgdata,_ := yaml.Marshal(cfg)
		log.Println("ping")
		err = ioutil.WriteFile(cfgpath,cfgdata,0666)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Cfg file created at",cfgpath)
		log.Fatal("Please check and restart, exiting...")

	}

	err = yaml.Unmarshal(cfgfile,&cfg)

	if err != nil {
		log.Fatal(err)
	}
	return
}
