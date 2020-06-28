package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Name string
	TemplatePath string
	Args []string
}

func Parse(filePath string) Config {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	config := Config{}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		panic(err)
	}
	return config
}
