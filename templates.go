package main

import (
	"embed"
	"fmt"
	"html/template"
)

//go:embed templates/*
var templates embed.FS
var tplCache map[string]*template.Template = make(map[string]*template.Template)

func loadTemplate(name string) *template.Template {
	tpl, ok := tplCache[name]
	if ok {
		return tpl
	}
	file, err := templates.ReadFile("templates/" + name)
	if err != nil {
		panic(fmt.Sprintf("Unable to load template %s", err.Error()))
	}
	tpl, err = template.New(name).Parse(string(file))
	if err != nil {
		panic(fmt.Sprintf("Unable to parse template %s", err.Error()))
	}
	tplCache[name] = tpl
	return tpl
}
