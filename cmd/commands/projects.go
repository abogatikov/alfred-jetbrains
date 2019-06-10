// Copyright 2019 Alex Bogatikov <abogatikov@devalexb.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package commands

import (
	"log"
)

// Projects retrieve project list
type Projects struct {
	Filter string `long:"filter" description:"Project name filter" default:""`
}

func (so *Projects) Execute([]string) error {
	wf.Run(so.run)
	return nil
}

func (so *Projects) run() {
	log.Printf("projects")
	defer func() {
		log.Printf("filter output: %s", so.Filter)
		if so.Filter != "" {
			wf.Filter(so.Filter)
		}
		log.Printf("return projects output")
		wf.SendFeedback()
	}()

	log.Printf("get recent projects")
	for _, item := range recentFolderItems() {
		log.Printf("add item: %s", item.Name)
		wf.NewItem(item.Name).
			Subtitle(item.Subtitle).
			UID(item.Name).
			Arg(item.Arg).
			Autocomplete(item.Name).
			Valid(true).
			Icon(createIcon(item.FolderName)).
			IsFile(false)
	}
}
