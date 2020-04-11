// Copyright 2019 Alex Bogatikov <abogatikov@devalexb.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type item struct {
	Name       string
	Subtitle   string
	Arg        string
	FolderName string
}

func recentFolderItems() []*item {
	names := make([]*item, 0)
	app, err := getApplication()

	if err != nil {
		log.Println(err)
		return names
	}

	for _, v := range app.Component.Option {
		if v.Name == "recentPaths" {
			for _, op := range v.List.Option {
				names = append(names, &item{
					Name:       folderName(op.Value),
					Subtitle:   projectDir(op.Value),
					Arg:        projectArg(op.Value),
					FolderName: appFolder(cApplication),
				})
			}

			return names
		}
	}

	return names
}

func appItems() []*item {
	names := make([]*item, 0)
	folders, err := ioutil.ReadDir(cToolboxFolder)

	if err != nil {
		log.Println(err)
		return names
	}

	for _, folder := range folders {
		names = append(names, &item{
			Name:       appName(folder.Name()),
			Subtitle:   appSubtitle(folder.Name()),
			Arg:        appArgString(folder.Name()),
			FolderName: folder.Name(),
		})
	}

	return names
}

func appArgString(appFolder string) string {
	appName := appName(appFolder)
	appNameLower := strings.ToLower(appName)
	appNameLowerReplacePlusesBuDashes := strings.ReplaceAll(appNameLower, " + ", "_")

	return strings.ReplaceAll(appNameLowerReplacePlusesBuDashes, " ", "_")
}

func appSubtitle(appFolder string) string {
	return fmt.Sprintf("%s Projects", appName(appFolder))
}

func projectArg(projectFolder string) string {
	return fmt.Sprintf("%s %s", cApplication, projectDir(projectFolder))
}

func appName(appFolder string) string {
	return strings.ReplaceAll(appFolder, ".app", "")
}

func folderName(l string) string {
	split := strings.Split(l, "/")
	// nolint:gomnd // this is a magic number
	if len(split) > 1 {
		return split[len(split)-1]
	}

	return l
}

func projectDir(projectFolder string) string {
	return strings.Replace(projectFolder, "$USER_HOME$", "~", 1)
}

func appFolder(name string) string {
	switch name {
	case "android_studio":
		return "Android Studio.app"
	case "appcode":
		return "AppCode.app"
	case "clion":
		return "CLion.app"
	case "datagrip":
		return "DataGrip.app"
	case "goland":
		return "GoLand.app"
	case "intellij_idea_community_educational":
		return "IntelliJ IDEA Community Educational.app"
	case "intellij_idea_community":
		return "IntelliJ IDEA Community.app"
	case "intellij_idea_ultimate_jbr11":
		return "IntelliJ IDEA Ultimate + JBR11.app"
	case "intellij_idea_ultimate":
		return "IntelliJ IDEA Ultimate.app"
	case "phpstorm":
		return "PhpStorm.app"
	case "pycharm_community":
		return "PyCharm Community.app"
	case "pycharm_edu":
		return "PyCharm Edu.app"
	case "pycharm_professional":
		return "PyCharm Professional.app"
	case "rider":
		return "Rider.app"
	case "rubymine":
		return "RubyMine.app"
	case "webstorm":
		return "WebStorm.app"
	default:
		return ""
	}
}
