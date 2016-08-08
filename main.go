package main

import (
  "net/http"
  "bytes"
  "html/template"
  "github.com/boltdb/bolt"
  "github.com/GeertJohan/go.rice"
  "flag"
)

type Config struct {
  Domain string
  SecretKey string
  ClientKey string
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
      link, storeErr := storeAndLink(db, body)
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
        secret, readErr := readAndDelete(db, uid)
        if readErr != nil {
          tpls["error.tpl"].Execute(&tplRes, nil)
        } else if (secret == "") {
          tpls["empty.tpl"].Execute(&tplRes, nil)
        } else {
          tpls["show.tpl"].Execute(&tplRes, secret)
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

  var port = flag.String("port", "8080", "port for server")
  var domain = flag.String("domain", "http://localhost:8080", "server domain name")
  var dbLocation = flag.String("db", "potemkin.db", "database file location")
  var secretKey = flag.String("secretKey", "", "Recaptcha secret key")
  var clientKey = flag.String("clientKey", "", "Recaptcha client key")
  flag.Parse()

  db, err := bolt.Open(*dbLocation, 0600, nil)
  defer db.Close()

  config := Config{*domain, *secretKey, *clientKey}

  if err != nil {
    panic(err)
  }

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    handler(w, r, tpls, db, staticBox, config)
  })
  http.ListenAndServe(":" + *port, nil)
}
