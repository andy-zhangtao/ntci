package main

const cloneTpl = `#!/bin/sh
cd {{.Root}}; git clone --branch {{.Branch}} "{{.Url}}" {{.Name}} 2>&1; cd {{.Name}}; git checkout -qf $NTCI_BUILDER_SHA; base64 -d $NTCI_BUILDER_DOCKERFILE > Dockerfile`

const buildTpl = `#!/bin/sh
echo "------->[Commit SHA]"
echo $NTCI_BUILDER_SHA
echo ""
git log -1 --pretty=%B |xargs echo "commit message: " 
echo "------->[Environment]"
echo ""
env |grep NTCI|grep -v NTCI_BUILDER_TOKEN|grep -v PASSWD|grep -v PASSWORD|grep -v password|grep -v passwd
echo ""
echo "------->[Dockerfile]"
cat Dockerfile
echo ""
echo "------->[Before Build]"
{{range .BeforeBuild}}
echo '  =====> {{.}}'
{{.}} 2>&1 ||true
{{end}}
echo ''

echo "------->[Build]"
set -e
{{range .Build}}
echo '  =====> {{.}}'
{{.}} 2>&1
{{end}}
echo ''

echo "------->[After Build]"
{{range .AfterBuild}}
echo '  =====> {{.}}'
{{.}} 2>&1 ||true
{{end}}
echo ''
`
