package main

import (
  "net/http"
  "bytes"
  "html/template"
  "github.com/boltdb/bolt"
  "github.com/GeertJohan/go.rice"
  "flag"
  "github.com/tkanos/gonfig"
)

type Config struct {
  Domain string
  SecretKey string
  ClientKey string
  DbLocation string
  Port string
}

func handler(w http.ResponseWriter, r *http.Request, tpls TplRepo, db *bolt.DB, box *rice.Box, config Config) {
  var action = r.URL.Path[1:];
  switch action {

  case "apple-touch-icon.png":
    fallthrough
  case "Roboto-Regular.woff2":
    fallthrough
  case "favicon.ico":
    fallthrough
  case "tile-wide.png":
    fallthrough
  case "tile.png":
    fallthrough
  case "main.js":
    fallthrough
  case "materialize.css":
    w.Header().Set("Content-Type", "text/css");
    fallthrough
  case "robots.txt":
    w.WriteHeader(200)
    file, _ := box.Bytes(action)
    w.Write(file)

  case "":
    var tplRes bytes.Buffer
    tpls["new.tpl"].Execute(&tplRes, nil)
    tpls["layout.tpl"].Execute(w, template.HTML(tplRes.String()))
  case "create":
    var tplRes bytes.Buffer
    err := r.ParseForm()
    if (err != nil) {
      tpls["error.tpl"].Execute(&tplRes, err)
    } else {
      body := r.FormValue("secret")
      salt := r.FormValue("salt")
      link, storeErr := storeAndLink(db, body, salt)
      link = config.Domain + "/show?uid=" + link
      if storeErr != nil {
        tpls["error.tpl"].Execute(&tplRes, nil)
      } else {
        tpls["create.tpl"].Execute(&tplRes, link)
      }
    }
    tpls["layout.tpl"].Execute(w, template.HTML(tplRes.String()))

  case "show":
    var tplRes bytes.Buffer
    err := r.ParseForm()
    if (err != nil) {
      tpls["error.tpl"].Execute(&tplRes, err)
    } else {
      uid := r.FormValue("uid")
      recaptcha := r.FormValue("g-recaptcha-response")

      if recaptcha != "" && checkRecaptcha(config.SecretKey, recaptcha) {
        secret, salt, readErr := readAndDelete(db, uid)
        if readErr != nil {
          tpls["error.tpl"].Execute(&tplRes, nil)
        } else if (secret == "") {
          tpls["empty.tpl"].Execute(&tplRes, nil)
        } else {
          tpls["show.tpl"].Execute(&tplRes, map[string]string{
            "secret": secret,
            "salt": salt,
          })
        }
      } else {
        tpls["preshow.tpl"].Execute(&tplRes, map[string]string{
          "uid": uid,
          "clientKey": config.ClientKey,
        })
      }
    }
    tpls["layout.tpl"].Execute(w, template.HTML(tplRes.String()))
  default:
    w.WriteHeader(http.StatusNotFound)
    tpls["404.html"].Execute(w, "")
  }
}

func main() {
  tpls := preloadTemplates()

  staticBox := rice.MustFindBox("static")
  var configLoc = flag.String("config", "./config.json", "config file for potemkin")
  flag.Parse()

  config := Config{}

  gonfig.GetConf(*configLoc, &config)

  db, err := bolt.Open(config.DbLocation, 0600, nil)
  defer func() {
    if db != nil {
      db.Close()
    }
  }()

  if err != nil {
    panic(err)
  }

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    handler(w, r, tpls, db, staticBox, config)
  })
  http.ListenAndServe(":" + config.Port, nil)
}
