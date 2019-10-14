package main

const cloneTpl = `#!/bin/sh
cd {{.Root}}; git clone --branch {{.Branch}} "{{.Url}}" {{.Name}} > /build.log 2>&1`