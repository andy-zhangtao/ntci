package main

const cloneTpl = `#!/bin/sh
cd {{.Root}}; git clone "{{.Url}}" {{.Name}} > /build.log 2>&1 ; cd {{.Root}}{{.Name}}; git checkout -b {{.Branch}} origin/{{.Branch}} >> /build.log 2>&1
`