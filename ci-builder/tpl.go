package main

//const cloneTpl = `#!/bin/sh
//cd {{.Root}}; git clone --branch {{.Branch}} "{{.Url}}" {{.Name}} > /build.log 2>&1`
const cloneTpl = `#!/bin/sh
cd {{.Root}}; git clone --branch {{.Branch}} "{{.Url}}" {{.Name}} 2>&1`

//const buildTpl = `#!/bin/sh
//{{range .Env}}
//echo 'set {{.}}'
//set {{.}}
//{{end}}
//
//echo "Environment"
//env >> /build.log
//echo "  "
//
//{{range .BeforeBuild}}
//echo '{{.}}'
//{{.}} >> /build.log 2>&1 ||true
//{{end}}
//
//set -e
//{{range .Build}}
//echo '{{.}}'
//{{.}} >> /build.log 2>&1
//{{end}}
//
//{{range .AfterBuild}}
//echo '{{.}}'
//{{.}} >> /build.log 2>&1 ||true
//{{end}}
//`

const buildTpl = `#!/bin/sh
{{range .Env}}
echo 'set {{.}}'
set {{.}}
{{end}}

echo "Environment"
env 
echo "  "

{{range .BeforeBuild}}
echo '{{.}}'
{{.}} ||true
{{end}}

set -e
{{range .Build}}
echo '{{.}}'
{{.}} 2>&1
{{end}}

{{range .AfterBuild}}
echo '{{.}}'
{{.}} ||true
{{end}}
`