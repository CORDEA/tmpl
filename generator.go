package main

import (
	"os"
	"path/filepath"
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
