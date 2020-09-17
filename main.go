package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Lmicroseconds | log.LUTC)
	var tplFile string
	var dataFile string
	flag.StringVar(&tplFile, "t", "", "path to template file")
	flag.StringVar(&dataFile, "d", "", "path to data(JSON) file(read from stdin if not specified)")
	flag.Parse()
	if tplFile == "" {
		handle(errors.New("template file wasn't provided"))
	}
	var dataFileReader *os.File
	if dataFile == "" {
		dataFileReader = os.Stdin
	} else {
		f, err := os.Open(dataFile)
		if err != nil {
			handle(err)
		}
		defer f.Close()
		dataFileReader = f
	}
	data := make(map[string]interface{})
	dec := json.NewDecoder(dataFileReader)
	if err := dec.Decode(&data); err != nil {
		handle(err)
	}
	b, err := ioutil.ReadFile(tplFile)
	if err != nil {
		handle(err)
	}
	tplFileContent := string(b)
	t := template.Must(template.New("tpl").Parse(tplFileContent))
	if err := t.Execute(os.Stdout, data); err != nil {
		handle(err)
	}
}

func handle(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
