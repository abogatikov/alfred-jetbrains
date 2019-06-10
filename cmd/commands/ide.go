// Copyright 2019 Alex Bogatikov <abogatikov@devalexb.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package commands

import (
	"log"
)

// IDE retrieve IDE list
type IDE struct {
	Filter string `long:"filter" description:"Project name filter" default:""`
}

func (so *IDE) Execute([]string) error {
	wf.Run(so.run)
	return nil
}

func (so *IDE) run() {
	log.Printf("ide")
	defer func() {
		if so.Filter != "" {
			log.Printf("filter output: %s", so.Filter)
			wf.Filter(so.Filter)
		}
		log.Printf("return ide output")
		wf.SendFeedback()
	}()

	log.Printf("get applications")
	for _, name := range appItems() {
		log.Printf("add item: %s", name.FolderName)
		wf.NewItem(name.Name).
			Subtitle(name.Subtitle).
			UID(name.Arg).
			Arg(name.Arg).
			Autocomplete(name.Arg).
			Icon(createIcon(name.FolderName)).
			Valid(true).
			IsFile(false)
	}
}
