package handler

import (
	"html/template"
)

var layout = "template/layout.tmpl"
var nav = "template/nav.tmpl"
var footer = "template/footer.tmpl"

var templates = map[string]*template.Template{}

func init() {
	templates["index"] = template.Must(template.ParseFiles(layout, nav, footer, "template/index.tmpl"))
	templates["scan"] = template.Must(template.ParseFiles(layout, nav, footer, "template/scan.tmpl"))
	templates["settings"] = template.Must(template.ParseFiles(layout, nav, footer, "template/settings.tmpl"))
}
