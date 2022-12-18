package main

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const googleApi string = "https://www.google.com/recaptcha/api/siteverify"

type GoogleResponse struct {
	Success  bool   `json:"success"`
	Hostname string `json:"hostname"`
}

func checkRecaptcha(secretKey string, response string) bool {
	resp, err := http.PostForm(googleApi, url.Values{
		"response": {response},
		"secret":   {secretKey},
	})
	if err != nil {
		return false
	}
	target := GoogleResponse{}
	json.NewDecoder(resp.Body).Decode(&target)
	return target.Success
}
