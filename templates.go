package main

import (
  "html/template"
  "github.com/GeertJohan/go.rice"
)

var templates []string = []string{"404.html", "layout.tpl", "create.tpl", "new.tpl", "error.tpl", "show.tpl", "empty.tpl"}

type TemplateResult struct {name string; t *template.Template}
type TplRepo map[string]*template.Template

func preloadTemplates() TplRepo {
  templateBox := rice.MustFindBox("templates")
  tplChan := make(chan TemplateResult)
  for _, tpl := range templates {
    go loadTemplate(tpl, templateBox, tplChan)
  }
  result := make(TplRepo)
  for i := 0; i < len(templates); i++ {
    tplResult := <- tplChan
    result[tplResult.name] = tplResult.t
  }
  return result
}

func loadTemplate(tpl string, box *rice.Box, tplChan chan TemplateResult) {
  data, err := box.Bytes(tpl);

  if err != nil {
    panic(err)
  }

  t := template.Must(template.New(tpl).Parse(string(data)))

  tplChan <- TemplateResult{tpl, t}
}
