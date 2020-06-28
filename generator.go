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
	paths := g.listFiles()
	g.replacePath(paths)
}

func (g *generator) listFiles() []string {
	var files []string
	err := filepath.Walk(g.templatePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

func (g *generator) replacePath(files []string) {
	for i := range files {
		file := files[i]
		replaced := ""
		for _, lst := range strings.Split(file, string(filepath.Separator)) {
			if strings.HasPrefix(lst, "+") && strings.HasSuffix(lst, "+") {
				key := lst[1 : len(lst)-1]
				val, ok := g.args[key]
				if !ok {
					log.Fatalln("Required argument does not exist.", key)
				}
				replaced = filepath.Join(replaced, filepath.FromSlash(val))
			} else {
				replaced = filepath.Join(replaced, lst)
			}
		}
		files[i] = replaced
	}
}
