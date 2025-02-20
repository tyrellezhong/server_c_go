package txttemplate

import (
	"os"
	"text/template"
)

const tmpl = `{{define "T1"}}Hello, {{.Name}} ! -- by tmpl{{end}}`

const condition = `
{{- if .IsAdmin}}
Hello, Admin {{.Name}}!
{{- else}}
Hello, User {{.Name}}!
{{- end}}

{{- range .Items}}
- {{.}}
{{- end}}
`

type Person struct {
	Name    string
	IsAdmin bool
	Items   []string
}

func ConditionTest() {
	t := template.Must(template.New("example").Parse(condition))
	p := Person{
		Name:    "Alice",
		IsAdmin: true,
		Items:   []string{"Item1", "Item2", "Item3"},
	}
	t.Execute(os.Stdout, p)
}
