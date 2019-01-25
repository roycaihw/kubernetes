/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"k8s.io/apiextensions-apiserver/pkg/controller/openapi"
)

var (
	boilerplateFile string
	outputFile      string
)

func main() {
	flag.Parse()
	if len(boilerplateFile) == 0 || len(outputFile) == 0 {
		log.Fatal("boilerplate and output path required")
	}
	b, err := ioutil.ReadFile(boilerplateFile)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.WriteString(fmt.Sprintf("%spackage openapi", string(b)))

	for _, t := range openapi.Templates {
		swagger := openapi.SwaggerTemplate(t.Scope, t.Status, t.Scale)
		swaggerJSON, err := json.MarshalIndent(swagger, " ", " ")
		if err != nil {
			panic(err)
		}
		f.WriteString(fmt.Sprintf("\n\nconst %s = []byte(```%s```)", t.Name, string(swaggerJSON)))
	}
	f.WriteString(fmt.Sprintf("\n"))
}

func init() {
	flag.StringVar(&boilerplateFile, "boilerplate-file", "", "Path to a boilerplate file.")
	flag.StringVar(&outputFile, "output-file", "", "Path to output file.")
}
