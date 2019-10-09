package main

import (
	"os"
	"strconv"

	"ntci/ci-agent/web"
)

/***
* ci-agents
* Listen quest from git repository. So there will has:
* + web server
*
 */
func main() {
	port := 8000

	p, err := strconv.Atoi(os.Getenv("CI_WEB_PORT"))
	if err == nil {
		port = p
	}

	web.Run(port)
}
