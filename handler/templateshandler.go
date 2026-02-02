package handler

import "html/template"

var templates *template.Template

func Temp(t *template.Template) {
	templates = t
}
