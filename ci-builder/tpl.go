package main

const cloneTpl = `#!/bin/sh
cd {{.Root}}; git clone --branch {{.Branch}} "{{.Url}}" {{.Name}} 2>&1`

const buildTpl = `#!/bin/sh
echo "------->Commit SHA"
echo $NTCI_BUILDER_SHA
echo ""
echo "------->Environment"
echo ""
env |grep -v NTCI_BUILDER_TOKEN
echo ""
echo "------->Before Build"
{{range .BeforeBuild}}
echo '-------> {{.}}'
{{.}} ||true
{{end}}
echo ''

echo "------->Build"
set -e
{{range .Build}}
echo '-------> {{.}}'
{{.}} 2>&1
{{end}}
echo ''

echo "------->After Build"
{{range .AfterBuild}}
echo '-------> {{.}}'
{{.}} ||true
{{end}}
echo ''
`
