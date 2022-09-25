package templates

import (
	_ "embed"
	"html/template"
)

var IndexView *template.Template
var ErrorView *template.Template

//go:embed index.html
var indexTemplate string

//go:embed error.html
var errorTemplate string

//go:embed user.html
var userViewBase string

//go:embed register.html
var registerBase string

//go:embed static/teststat.html
var dateConv string

//go:embed base.html
var baseTemplate string

func init() {
	ErrorView = getTemplate(errorTemplate)
	IndexView = getTemplate(indexTemplate)
}

func getTemplate(pageTemplate string) *template.Template {
	base := template.Must(template.New("base").Parse(baseTemplate))
	built := template.Must(base.Parse(pageTemplate))
	return built
}
