package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	TemplatePath string   `yaml:"templatePath"`
	Args         []string `yaml:"args"`
}

func Parse(filePath string) Config {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	config := Config{}
	if err = yaml.Unmarshal(file, &config); err != nil {
		panic(err)
	}
	return config
}
