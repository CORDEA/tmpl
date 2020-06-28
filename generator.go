package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type generator struct {
	templatePath string
	args         map[string]string
}

func (g *generator) generate() error {
	files, err := g.listFiles()
	if err != nil {
		return err
	}
	for _, path := range files {
		target, err := g.makeTargetPath(path)
		if err != nil {
			return err
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		tmpl, err := template.New("").Parse(string(data))
		if err != nil {
			return err
		}
		if err = os.MkdirAll(filepath.Dir(target), 0755); err != nil {
			return err
		}
		file, err := os.Create(target)
		if err != nil {
			return err
		}
		if err = tmpl.Execute(file, g.args); err != nil {
			return err
		}
		if err = file.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (g *generator) listFiles() ([]string, error) {
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
	return files, err
}

func (g *generator) makeTargetPath(path string) (string, error) {
	path, err := filepath.Rel(g.templatePath, path)
	if err != nil {
		return "", err
	}
	newPath := ""
	for _, lst := range strings.Split(path, string(filepath.Separator)) {
		if strings.HasPrefix(lst, "+") && strings.HasSuffix(lst, "+") {
			key := lst[1 : len(lst)-1]
			val, ok := g.args[key]
			if !ok {
				return "", errors.New("Required argument does not exist. " + key)
			}
			newPath = filepath.Join(newPath, filepath.FromSlash(val))
		} else {
			newPath = filepath.Join(newPath, lst)
		}
	}
	return newPath, nil
}
