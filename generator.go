package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type generator struct {
	templatePath string
	args         map[string]string
}

func (g *generator) generate() {
	g.listFiles()
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
