package main

const cloneTpl = `#!/bin/sh
cd {{.Root}}; git clone --branch {{.Branch}} "{{.Url}}" {{.Name}} 2>&1`

const buildTpl = `#!/bin/sh
echo "------->Environment"
echo ""
env |grep -v NTCI_BUILDER_TOKEN
echo ""

{{range .BeforeBuild}}
echo '-------> {{.}}'
{{.}} ||true
{{end}}
echo ''

set -e
{{range .Build}}
echo '-------> {{.}}'
{{.}} 2>&1
{{end}}
echo ''

{{range .AfterBuild}}
echo '-------> {{.}}'
{{.}} ||true
{{end}}
echo ''
`