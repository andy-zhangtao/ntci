package main

import "ntci/ci-agent/web"

/***
* ci-agents
* Listen quest from git repository. So there will has:
* + web server
*
 */
func main() {

	web.Run(8000)
}
