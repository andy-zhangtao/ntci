package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var jid string
var buildAddr string
var url string
var branch string
var token string

type gitMeta struct {
	Name   string
	Root   string
	Url    string
	Branch string
	Token  string
}

var gm gitMeta

func init() {

	logrus.SetLevel(logrus.DebugLevel)

	if os.Getenv("NTCI_BUILDER_JID") != "" {
		jid = os.Getenv("NTCI_BUILDER_JID")
	} else {
		logrus.Fatalf("NTCI_BUILDER_JID EMPTY!")
	}

	if os.Getenv("NTCI_BUILDER_ADDR") != "" {
		buildAddr = os.Getenv("NTCI_BUILDER_ADDR")
	} else {
		logrus.Fatalf("NTCI_BUILDER_ADDR EMPTY!")
	}

	if os.Getenv("NTCI_BUILDER_GIT") != "" {
		url = os.Getenv("NTCI_BUILDER_GIT")
	} else {
		logrus.Fatalf("NTCI_BUILDER_GIT EMPTY!")
	}

	if os.Getenv("NTCI_BUILDER_BRANCH") != "" {
		branch = os.Getenv("NTCI_BUILDER_BRANCH")
	} else {
		logrus.Fatalf("NTCI_BUILDER_BRANCH EMPTY!")
	}

	r := "~"
	if os.Getenv("NTCI_BUILDER_ROOT") != "" {
		r = os.Getenv("NTCI_BUILDER_ROOT")
	}

	if os.Getenv("NTCI_BUILDER_TOKEN") != "" {
		token = os.Getenv("NTCI_BUILDER_TOKEN")
	}

	if token != "" {
		if strings.HasPrefix(url, "http://") {
			url = fmt.Sprintf("http://%s@%s", token, strings.Split(url, "http://")[1])
		} else {
			url = fmt.Sprintf("https://%s@%s", token, strings.Split(url, "https://")[1])
		}
	}

	gm = gitMeta{
		Name:   jid,
		Root:   r,
		Url:    url,
		Branch: branch,
		Token:  token,
	}

	debugGM()
}

func debugGM() {
	logrus.Debug("=============================")
	logrus.Debugf("Name: %s", gm.Name)
	logrus.Debugf("Branch: %s", gm.Branch)
	logrus.Debugf("Root: %s", gm.Root)
	logrus.Debugf("Token: %s", gm.Token)
	logrus.Debugf("Url: %s", gm.Url)
	logrus.Debug("=============================")
}
