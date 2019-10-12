package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

var jid string
var buildAddr string
var url string
var branch string

func init() {
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
}
