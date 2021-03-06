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
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"syscall"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := path.Dir(ex)

	wd, _ := os.Getwd()
	_ = wd

	ctx, cancel := context.WithCancel(context.Background())

	cmd := exec.CommandContext(ctx, path.Join(dir, fmt.Sprintf("beagle-%s", os.Args[1])), os.Args[2:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt)
		signal.Notify(signalChan, syscall.SIGTERM)

		select {
		case s := <-signalChan:
			cmd.Process.Signal(s)
		}
	}()

	_ = cancel

	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
