package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"ntci/ci-agent/web"
)

var port = 8000

/***
* ci-agents
* Listen quest from git repository. So there will has:
* + web server
*
 */
func main() {
	web.Run(port)
}

func init() {
	p, err := strconv.Atoi(os.Getenv("CI_WEB_PORT"))
	if err == nil {
		port = p
	}

	switch strings.ToLower(os.Getenv("CI_WEB_LOG_LEVEL")) {
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
	}
}
