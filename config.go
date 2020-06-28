package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type config struct {
	name string
	templatePath string
	args []string
}

func parse(filePath string) config {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	config := config{}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		panic(err)
	}
	return config
}
