package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/elos/metis"
	"github.com/elos/metis/templates"
)

var RouterTemplate templates.Name = 0

func main() {
	models, _ := metis.ParseGlob("./definitions/models/*json")
	s := metis.BuildSchema(models...)

	e := templates.NewEngine("./", &templates.TemplateSet{
		RouterTemplate: []string{"router.tmpl"},
	})
	err := e.ParseTemplates()
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	if err := e.Execute(&buf, RouterTemplate, s); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("../router.go", buf.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}
	if err := exec.Command("goimports", "-w=true", "../router.go").Run(); err != nil {
		log.Fatal(err)
	}
}
