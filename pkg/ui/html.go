package ui

import (
	template2 "html/template"
	"os"
)

func htmlPath() string {
	wd, _ := os.Getwd()
	return wd + "/pkg/ui/dashboard.html"
}

func template() *template2.Template {
	return template2.Must(template2.ParseFiles(htmlPath()))
}

