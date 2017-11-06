/*
Copyright 2017 The Kubernetes Authors.

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
	"fmt"
	"os"

	"k8s.io/code-generator/cmd/api-linter/linters"
	"k8s.io/gengo/args"

	"github.com/golang/glog"
	"github.com/spf13/pflag"
)

const helpText = `# API linter

The API linter follows the same rules as openapi-gen, and validates API conventions on every type or package
that requires API definition generation. The following rules are the same from openapi-gen:

- To generate definition for a specific type or package add "+k8s:openapi-gen=true" tag to the type/package comment lines.
- To exclude a type or a member from a tagged package/type, add "+k8s:openapi-gen=false" tag to the comment lines.
`

func main() {
	arguments := args.Default()
	arguments.OutputFileBaseName = "api_linter_dummy"

	customArgs := &linters.CustomArgs{
		WhitelistFilename: "",
	}
	pflag.CommandLine.StringVar(&customArgs.WhitelistFilename, "whitelist-file", customArgs.WhitelistFilename, "Specify a csv file that contains the whitelist.")
	arguments.CustomArgs = customArgs

	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s\n", helpText)
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		pflag.PrintDefaults()
	}

	// Run it.
	if err := arguments.Execute(
		linters.NameSystems(),
		linters.DefaultNameSystem(),
		linters.Packages,
	); err != nil {
		glog.Fatalf("Error: %v", err)
	}
}
