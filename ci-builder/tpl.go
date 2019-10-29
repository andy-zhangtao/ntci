package main

const cloneTpl = `#!/bin/sh
cd {{.Root}}; git clone --branch {{.Branch}} "{{.Url}}" {{.Name}} 2>&1; cd {{.Name}}; git checkout -qf $NTCI_BUILDER_SHA`

const buildTpl = `#!/bin/sh
echo "------->Commit SHA"
echo $NTCI_BUILDER_SHA
echo ""
echo "------->Environment"
echo ""
env |grep NTCI|grep -v NTCI_BUILDER_TOKEN|grep -v PASSWD|grep -v PASSWORD|grep -v password|grep -v passwd
echo ""
echo "------->Before Build"
{{range .BeforeBuild}}
echo '  =====> {{.}}'
{{.}} 2>&1 ||true
{{end}}
echo ''

echo "------->Build"
set -e
{{range .Build}}
echo '  =====> {{.}}'
{{.}} 2>&1
{{end}}
echo ''

echo "------->After Build"
{{range .AfterBuild}}
echo '  =====> {{.}}'
{{.}} 2>&1 ||true
{{end}}
echo ''
`
