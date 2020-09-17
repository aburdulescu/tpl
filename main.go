package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

func main() {
	var tplFile string
	var dataFile string
	flag.StringVar(&tplFile, "t", "", "path to template file")
	flag.StringVar(&dataFile, "d", "", "path to data(JSON) file")
	flag.Parse()
	if tplFile == "" {
		handle("template file wasn't provided")
	}
	if tplFile == "" {
		handle("data file wasn't provided")
	}
	b, err := ioutil.ReadFile(tplFile)
	if err != nil {
		handle(err.Error())
	}
	tplFileContent := string(b)
	t := template.Must(template.New("makefile").Parse(tplFileContent))
	b, err = ioutil.ReadFile(dataFile)
	if err != nil {
		handle(err.Error())
	}
	data := make(map[string]interface{})
	if err := json.Unmarshal(b, &data); err != nil {
		handle(err.Error())
	}
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Println(err)
		handle(err.Error())
	}
}

func handle(err string) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
