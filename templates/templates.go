package templates

import (
	_ "embed"
	"html/template"
	"reflect"
)

var IndexView *template.Template
var ErrorView *template.Template
var RegisterView *template.Template
var DashboardView *template.Template
var LoginView *template.Template
var DebugView *template.Template
var CommentsView *template.Template

//go:embed index.html
var indexTemplate string

//go:embed error.html
var errorTemplate string

//go:embed dashboard.html
var dashboardTemplate string

//go:embed login.html
var loginTemplate string

//go:embed debug.html
var debugTemplate string

//go:embed user.html
var userViewBase string

//go:embed comments.html
var commentsTemplate string

//go:embed register.html
var registerTemplate string

//go:embed base.html
var baseTemplate string

func init() {
	ErrorView = getTemplate(errorTemplate)
	IndexView = getTemplate(indexTemplate)
	RegisterView = getTemplate(registerTemplate)
	DashboardView = getTemplate(dashboardTemplate)
	LoginView = getTemplate(loginTemplate)
	DebugView = getTemplate(debugTemplate)
	CommentsView = getTemplate(commentsTemplate)

}

func getTemplate(pageTemplate string) *template.Template {
	base := template.Must(template.New("base").Funcs(template.FuncMap{"hasField": hasField}).Parse(baseTemplate))
	built := template.Must(base.Parse(pageTemplate))
	return built
}

// hasField is a function called in templates to see if a given . has a given field. Credit:https://stackoverflow.com/questions/34703133/field-detection-in-go-html-template
func hasField(v interface{}, name string) bool {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return false
	}
	return rv.FieldByName(name).IsValid()
}
