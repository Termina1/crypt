package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/skip2/go-qrcode"
)

//go:embed static/*
var static embed.FS

type Config struct {
	Domain     string `required:"true" json:"domain"`
	SecretKey  string `required:"true" json:"secret_key"`
	ClientKey  string `required:"true" json:"client_key"`
	DbLocation string `default:"crypt.db" json:"db_location"`
	Port       string `default:"8080" json:"port"`
}

func handler(w http.ResponseWriter, r *http.Request, db *bolt.DB, config Config) {
	var action = r.URL.Path[1:]
	switch action {
	case "":
		var tplRes bytes.Buffer
		loadTemplate("new.tpl").Execute(&tplRes, nil)
		loadTemplate("layout.tpl").Execute(w, template.HTML(tplRes.String()))
	case "create":
		var tplRes bytes.Buffer
		err := r.ParseForm()
		if err != nil {
			loadTemplate("error.tpl").Execute(&tplRes, err)
		} else {
			body := r.FormValue("secret")
			salt := r.FormValue("salt")
			link, storeErr := storeAndLink(db, body, salt)
			link = config.Domain + "/show?uid=" + link
			if storeErr != nil {
				loadTemplate("error.tpl").Execute(&tplRes, nil)
			} else {
				loadTemplate("create.tpl").Execute(&tplRes, link)
			}
		}
		loadTemplate("layout.tpl").Execute(w, template.HTML(tplRes.String()))

	case "show":
		var tplRes bytes.Buffer
		err := r.ParseForm()
		if err != nil {
			loadTemplate("error.tpl").Execute(&tplRes, err)
		} else {
			uid := r.FormValue("uid")
			recaptcha := r.FormValue("g-recaptcha-response")

			if recaptcha != "" && checkRecaptcha(config.SecretKey, recaptcha) {
				secret, salt, readErr := readAndDelete(db, uid)
				if readErr != nil {
					loadTemplate("error.tpl").Execute(&tplRes, nil)
				} else if secret == "" {
					loadTemplate("empty.tpl").Execute(&tplRes, nil)
				} else {
					loadTemplate("show.tpl").Execute(&tplRes, map[string]string{
						"secret": secret,
						"salt":   salt,
					})
				}
			} else {
				loadTemplate("preshow.tpl").Execute(&tplRes, map[string]string{
					"uid":       uid,
					"clientKey": config.ClientKey,
				})
			}
		}
		loadTemplate("layout.tpl").Execute(w, template.HTML(tplRes.String()))
	case "qr.png":
		uid := r.URL.Query().Get("uid")
		data, err := qrcode.Encode(fmt.Sprintf("%s/show?uid=%s", config.Domain, uid), qrcode.Medium, 256)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "image/png")
			w.Write([]byte(""))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		}
	default:
		file, err := static.ReadFile("static/" + action)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			loadTemplate("404.html").Execute(w, "")
		} else {
			if strings.HasSuffix(action, ".css") {
				w.Header().Set("Content-Type", "text/css")
			}
			w.WriteHeader(http.StatusOK)
			w.Write(file)
		}
	}
}

func getEnv(prefix, name string) (string, bool) {
	envName := strings.ToUpper(fmt.Sprintf("%s_%s", prefix, name))
	return os.LookupEnv(envName)
}

func overrideConfig(prefix string, config any) error {
	type_ := reflect.TypeOf(config)
	if type_.Kind() != reflect.Ptr {
		return fmt.Errorf("config struct must be pointer to a struct")
	}
	type_ = type_.Elem()
	if type_.Kind() != reflect.Struct {
		return fmt.Errorf("config struct must be pointer to a struct")
	}
	value_ := reflect.ValueOf(config).Elem()
	fields := reflect.VisibleFields(type_)
	for _, field := range fields {
		if field.Type.Kind() != reflect.String {
			return fmt.Errorf("only string config parameters are supported, for %s given %s", field.Name, field.Type.String())
		}
		v, ok := getEnv(prefix, field.Name)
		fieldValue := value_.FieldByIndex(field.Index)
		if ok && fieldValue.CanSet() {
			fieldValue.SetString(v)
		}
		required, ok := field.Tag.Lookup("required")
		if ok && required == "true" && fieldValue.String() == "" {
			return fmt.Errorf("field %s is required but not set", field.Name)
		}
		def, ok := field.Tag.Lookup("default")
		if ok && fieldValue.String() == "" && fieldValue.CanSet() {
			fieldValue.SetString(def)
		}
	}
	return nil
}

func main() {
	config := Config{}
	var configLoc = flag.String("config", "", "config file for crypt in JSON format")
	var prefix = flag.String("prefix", "crypt", "env prefix for overrides")
	flag.Parse()
	if configLoc != nil && *configLoc != "" {
		file, err := os.ReadFile(*configLoc)
		if err != nil {
			panic(fmt.Sprintf("can't read config file %s", *configLoc))
		}
		json.Unmarshal(file, &config)
	}
	err := overrideConfig(*prefix, &config)
	if err != nil {
		panic(fmt.Sprintf("error while parsing config: %s", err.Error()))
	}
	db, err := bolt.Open(config.DbLocation, 0600, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, db, config)
	})
	fmt.Printf("Listening on port %s\n", config.Port)
	http.ListenAndServe(":"+config.Port, nil)
}
