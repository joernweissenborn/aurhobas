package config

type Config struct {
	Root string
}

var CFG Config

func StaticConfig(cfg Config) {
	CFG = cfg
}

func FromYaml(path string){}