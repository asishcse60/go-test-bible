package yamltohtml

import (
	"bytes"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
)

type PageData struct {
	Title string `yaml:"Title"`
	Desc  string `yaml:"Desc"`
	Body  string `yaml:"Body"`
}

func YamlToHtml(path string) (string, error) {
	tmpl, err := template.New("page").Parse(`<html><head><title>{{.Title}}</title></head><body>{{.Body}}</body></html>`)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	var pageData PageData

	err = yaml.Unmarshal(data, &pageData)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, pageData); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
