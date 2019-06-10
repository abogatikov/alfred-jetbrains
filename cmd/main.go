// Copyright 2019 Alex Bogatikov <abogatikov@devalexb.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/abogatikov/alfred-jetbrains/cmd/commands"
	"github.com/jessevdk/go-flags"
)

type Opts struct {
	IDE      commands.IDE      `command:"ide" description:"List of supported IDE"`
	Projects commands.Projects `command:"projects" description:"List of projects"`
}

func main() {
	var opts Opts
	parser := flags.NewParser(&opts, flags.Default)
	if _, err := parser.ParseArgs(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
