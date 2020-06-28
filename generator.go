package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type generator struct {
	templatePath string
	args         map[string]string
}

func (g *generator) generate() {
	for _, path := range g.listFiles() {
		target := g.makeTargetPath(path)
		data, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		tmpl, err := template.New("").Parse(string(data))
		if err != nil {
			panic(err)
		}
		if err = os.MkdirAll(filepath.Dir(target), 0755); err != nil {
			panic(err)
		}
		file, err := os.Create(target)
		if err != nil {
			panic(err)
		}
		if err = tmpl.Execute(file, g.args); err != nil {
			panic(err)
		}
		if err = file.Close(); err != nil {
			panic(err)
		}
	}
}

func (g *generator) listFiles() []string {
	var files []string
	err := filepath.Walk(g.templatePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

func (g *generator) makeTargetPath(path string) string {
	path, err := filepath.Rel(g.templatePath, path)
	if err != nil {
		panic(err)
	}
	newPath := ""
	for _, lst := range strings.Split(path, string(filepath.Separator)) {
		if strings.HasPrefix(lst, "+") && strings.HasSuffix(lst, "+") {
			key := lst[1 : len(lst)-1]
			val, ok := g.args[key]
			if !ok {
				log.Fatalln("Required argument does not exist.", key)
			}
			newPath = filepath.Join(newPath, filepath.FromSlash(val))
		} else {
			newPath = filepath.Join(newPath, lst)
		}
	}
	return newPath
}
