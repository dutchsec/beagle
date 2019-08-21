// Copyright 2019 The DutchSec Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var ()

// Usage is a replacement usage function for the flags package.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\tbeagle license {path}\n")
	fmt.Fprintf(os.Stderr, "\thttps://github.com/beagleorg/beagle\n")
	flag.PrintDefaults()
}

var (
	Exclude = func(excludePaths []string) func(string) bool {
		return func(s string) bool {
			for _, excludePath := range excludePaths {
				if s == excludePath {
					return true
				}

			}
			return false
		}
	}([]string{
		".git",
		"vendor",
	})
)

func main() {
	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("%s: ", os.Args[0]))
	flag.Usage = Usage
	flag.Parse()

	// We accept either one directory or a list of files. Which do we have?
	args := flag.Args()
	if len(args) == 0 {
		// Default: process whole package in current directory.
		args = []string{"."}
	}

	data, err := ioutil.ReadFile("LICENSE")
	if err != nil {
		log.Fatalf("Could not read license file: %s", err.Error())
	}

	for _, arg := range args {
		filepath.Walk(arg, func(p string, info os.FileInfo, err error) error {
			if Exclude(path.Base(p)) {
				return filepath.SkipDir
			}

			if path.Ext(p) == ".go" {
			} else {
				return nil
			}

			// find package main
			// check if license already
			// have option to override license

			originalFile, err := ioutil.ReadFile(p)
			if err != nil {
				log.Fatalf("Could not read file: %s: %s", p, err.Error())
			}

			// find index to "package"
			reader := bufio.NewReader(bytes.NewReader(originalFile))
			scanner := bufio.NewScanner(reader)

			scanner.Split(bufio.ScanLines)

			output := &bytes.Buffer{}

			packageFound := false
			for scanner.Scan() {
				line := scanner.Text()
				if packageFound {
				} else if strings.HasPrefix(line, "package") {
					packageFound = true

					output.Write(data)
				} else {
					continue
				}

				fmt.Fprintln(output, line)
			}

			if err := ioutil.WriteFile(p, output.Bytes(), info.Mode()); err != nil {
				log.Fatalf("Could not write file: %s: %s", p, err.Error())
			}

			// originalFile
			// readall
			// writeall

			// output license
			_ = data

			// copy file
			// should we check?

			return nil
		})
	}
}
