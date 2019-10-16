package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var jname string
var jid string
var buildAddr string
var url string
var branch string
var token string

type gitMeta struct {
	Name   string
	Id     string
	Root   string
	Url    string
	Branch string
	Token  string
	User   string
}

var gm gitMeta

func init() {

	logrus.SetLevel(logrus.DebugLevel)
	gm = gitMeta{
		//Name:   jname,
		//Id:     jid,
		//Root:   r,
		//Url:    url,
		//Branch: branch,
		//Token:  token,
	}

	if os.Getenv("NTCI_BUILDER_USER") != "" {
		gm.User = os.Getenv("NTCI_BUILDER_USER")
	} else {
		logrus.Fatalf("NTCI_BUILDER_USER EMPTY!")
	}

	if os.Getenv("NTCI_BUILDER_JID") != "" {
		gm.Name = os.Getenv("NTCI_BUILDER_JID")
	} else {
		logrus.Fatalf("NTCI_BUILDER_JID EMPTY!")
	}

	if os.Getenv("NTCI_BUILDER_ID") != "" {
		gm.Id = os.Getenv("NTCI_BUILDER_ID")
	} else {
		logrus.Fatalf("NTCI_BUILDER_ID EMPTY!")
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
		gm.Branch = os.Getenv("NTCI_BUILDER_BRANCH")
	} else {
		logrus.Fatalf("NTCI_BUILDER_BRANCH EMPTY!")
	}

	gm.Root = "~"
	if os.Getenv("NTCI_BUILDER_ROOT") != "" {
		gm.Root = os.Getenv("NTCI_BUILDER_ROOT")
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

	gm.Token = token
	gm.Url = url
	debugGM()
}

func debugGM() {
	logrus.Debug("=============================")
	logrus.Debugf("Build Name: %s", gm.Name)
	logrus.Debugf("Build ID: %s", gm.Id)
	logrus.Debugf("Build User: %s", gm.User)
	logrus.Debugf("Branch: %s", gm.Branch)
	logrus.Debugf("Root: %s", gm.Root)
	logrus.Debugf("Token: %s", gm.Token)
	logrus.Debugf("Url: %s", gm.Url)
	logrus.Debug("=============================")
}
