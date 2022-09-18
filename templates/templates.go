package templates

import (
	_ "embed"
	"html/template"
)

var UserViewSingle *template.Template
var RegistrationView *template.Template

//go:embed base.html
var baseTemplate string

//go:embed user.html
var userViewBase string

//go:embed register.html
var registerBase string

//go:embed static/teststat.html
var dateConv string

func init() {
	base := template.Must(template.New("base").Parse(baseTemplate))
	userpartial := template.Must(base.Parse(userViewBase))
	UserViewSingle = template.Must(userpartial.Parse(dateConv))
	base = template.Must(template.New("base").Parse(baseTemplate))
	RegistrationView = template.Must(base.Parse(registerBase))

}
