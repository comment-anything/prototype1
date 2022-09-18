package templates

import (
	_ "embed"
	"html/template"
)

var UserViewSingle *template.Template

//go:embed base.html
var baseTemplate string

//go:embed user.html
var userViewBase string

//go:embed static/teststat.html
var dateConv string

func init() {
	UserViewSingle = template.Must(template.Must(template.ParseFiles("templates/user.html")).Parse(dateConv))
	base := template.Must(template.New("base").Parse(baseTemplate))
	userpartial := template.Must(base.Parse(userViewBase))
	UserViewSingle = template.Must(userpartial.Parse(dateConv))

}
