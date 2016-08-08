package main

import (
  "net/http"
  "net/url"
  "encoding/json"
)

const googleApi string = "https://www.google.com/recaptcha/api/siteverify"

type GoogleResponse struct {
  Success bool
  Challenge_ts int
  Hostname string
}

func checkRecaptcha(response string) (bool) {
  secretKey := "6LeUEycTAAAAAFIUyamcvK6FSo5rCRwJEzRnnvLY"
  resp, err := http.PostForm(googleApi, url.Values{
    "response": { response },
    "secret": { secretKey },
  })

  if err != nil {
    return false
  }

  target := new(GoogleResponse)

  json.NewDecoder(resp.Body).Decode(target)

  return target.Success
}
