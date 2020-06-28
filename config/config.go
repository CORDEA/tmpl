package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strings"
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

func (c *Config) MapArgs(name string, args string) map[string]string {
	m := map[string]string{}
	m["name"] = name
	if args == "" {
		return m
	}
	for _, arg := range strings.Split(args, ",") {
		pair := strings.Split(arg, ":")
		if len(pair) != 2 {
			log.Fatalln("Illegal argument.", pair)
		}
		m[pair[0]] = pair[1]
	}
	if len(m) != len(c.Args) {
		log.Fatalln("Number of arguments do not match.")
	}
	for _, arg := range c.Args {
		if _, ok := m[arg]; !ok {
			log.Fatalln("Illegal argument.", arg)
		}
	}
	return m
}
