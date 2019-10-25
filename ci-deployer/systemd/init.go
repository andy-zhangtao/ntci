package main

import (
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/coreos/go-systemd/dbus"
	"github.com/sirupsen/logrus"
	"ntci/ci-deployer/model/config"
)

var c config.Configure
var conn *dbus.Conn

func init() {
	var err error

	switch strings.ToLower(os.Getenv("NTCI_DEPLOY_LOG_LEVEL")) {
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
	}

	if os.Getenv("NTCI_DEPLOY_CONFIG") != "" {
		if _, err := toml.DecodeFile(os.Getenv("NTCI_DEPLOY_CONFIG"), &c); err != nil {
			logrus.Fatal("Configure Init Failed! ")
		}
	}

	conn, err = dbus.New()
	if err != nil {
		logrus.Fatalf("Connect Systemd Error: %s", err.Error())
	}

	if c.Port == 0 {
		c.Port = 80
	}
	logrus.Debug("-------Configure-------")
	logrus.Debugf("Listen On: %d", c.Port)
	for key, value := range c.Units {
		logrus.Debugf("Name: %s Unit: %s", key, value)
	}
	logrus.Debug("-----------------------")

}
