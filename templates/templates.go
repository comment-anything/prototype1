package templates

import (
	_ "embed"
	"html/template"
)

var IndexView *template.Template
var ErrorView *template.Template
var RegisterView *template.Template
var DashboardView *template.Template
var LoginView *template.Template

//go:embed index.html
var indexTemplate string

//go:embed error.html
var errorTemplate string

//go:embed dashboard.html
var dashboardTemplate string

//go:embed login.html
var loginTemplate string

//go:embed user.html
var userViewBase string

//go:embed register.html
var registerTemplate string

//go:embed static/teststat.html
var dateConv string

//go:embed base.html
var baseTemplate string

func init() {
	ErrorView = getTemplate(errorTemplate)
	IndexView = getTemplate(indexTemplate)
	RegisterView = getTemplate(registerTemplate)
	DashboardView = getTemplate(dashboardTemplate)
	LoginView = getTemplate(loginTemplate)
}

func getTemplate(pageTemplate string) *template.Template {
	base := template.Must(template.New("base").Parse(baseTemplate))
	built := template.Must(base.Parse(pageTemplate))
	return built
}
