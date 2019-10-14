package main

const cloneTpl = `#!/bin/sh
cd {{.Root}}; git clone --branch {{.Branch}} "{{.Url}}" {{.Name}} > /build.log 2>&1`

const buildTpl = `#!/bin/sh
{{range .Env}}
echo 'set {{.}}'
set {{.}}
{{end}}
{{range .BeforeBuild}}
echo '{{.}}'
sh "{{.}}" >> /build.log 2>&1 ||true
{{end}}
{{range .Build}}
echo '{{.}}'
sh "{{.}}" >> /build.log 2>&1
{{end}}
{{range .AfterBuild}}
echo '{{.}}'
sh "{{.}}" >> /build.log 2>&1 ||true
{{end}}
`
