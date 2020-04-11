// Copyright 2019 Alex Bogatikov <abogatikov@devalexb.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package commands

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
)

type application struct {
	Component component `xml:"component"`
}

type component struct {
	Name   string    `xml:"name,attr"`
	Option []*option `xml:"option"`
}

type option struct {
	Name  string     `xml:"name,attr"`
	Value string     `xml:"value,attr"`
	List  optionList `xml:"list"`
}

type optionList struct {
	Option []*option `xml:"option"`
}

func getApplication() (*application, error) {
	b, err := ioutil.ReadFile(cApplicationRecentFolderFile)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var app application

	decoder := xml.NewDecoder(bytes.NewReader(b))
	err = decoder.Decode(&app)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &app, nil
}
