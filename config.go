package main

import (
	"os"
	"path"
	"strings"

	ini "gopkg.in/ini.v1"
)

type Config struct {
	RootDirectory string
	Lowercase     bool
}

var config Config

const defaultRootDirectory = "/tmp"

func init() {
	loadGitGlobalConfig()
}

func loadGitGlobalConfig() {
	filepath := path.Join(os.Getenv("HOME"), ".gitconfig")
	c, err := ini.Load(filepath)
	if err != nil {
		// TODO log
		return
	}

	section := c.Section("get")

	config.RootDirectory = section.Key("root").Validate(func(path string) string {
		if len(path) == 0 {
			return defaultRootDirectory
		}
		return interpretHomeTilde(path)
	})
	config.Lowercase = section.Key("lowercase").MustBool(false)
}

func interpretHomeTilde(path string) string {
	return strings.Replace(path, "~", os.Getenv("HOME"), 1)
}
