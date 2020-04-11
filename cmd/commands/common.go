// Copyright 2019 Alex Bogatikov <abogatikov@devalexb.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package commands

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	aw "github.com/deanishe/awgo"
)

const (
	iconSubPath         = "/Contents/Resources/icon.icns"
	cRecentProjectsPath = `/config/options/recentProjects.xml`
)

var (
	cConfigFolder                string
	cApplication                 string
	cToolboxFolder               string
	cApplicationRecentFolderFile string
	wf                           *aw.Workflow
)

func init() {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
		return
	}

	cToolboxFolder = os.Getenv("toolbox_application_links_folder")
	if strings.HasPrefix(cToolboxFolder, "~") {
		cToolboxFolder = filepath.Join(dir, strings.Replace(cToolboxFolder, "~", "", 1))
	}

	cConfigFolder = os.Getenv("config_folder")
	cApplication = os.Getenv("application")

	cConfigRecentFiles := fmt.Sprintf(path.Join(cConfigFolder, "/%s/", cRecentProjectsPath), cApplication)
	cApplicationRecentFolderFile = strings.Replace(cConfigRecentFiles, "~", dir, 1)

	wf = aw.New()
}

func createIcon(appFolder string) *aw.Icon {
	return &aw.Icon{
		Value: fmt.Sprintf(path.Join(cToolboxFolder, "/%s/", iconSubPath), appFolder),
	}
}
