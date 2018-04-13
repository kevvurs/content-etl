package main

import (
        "io/ioutil"
        "net/http"
        "os"
        "golang.org/x/net/context"
        "google.golang.org/appengine/log"
        "google.golang.org/appengine/urlfetch"
)

const API_URL = "https://newsapi.org/v2/top-headlines?country=us"

func fetchNews(ctx context.Context) []byte {
  secret := os.Getenv("APP_SECRET")
  if secret == "" {
    log.Errorf(ctx, "Failed to load news API credentials")
    return make([]byte, 0)
  }
  client := urlfetch.Client(ctx)
  req, err := http.NewRequest("GET", API_URL, nil)
  if err != nil {
    log.Errorf(ctx, "Failed to build API request")
    return make([]byte, 0)
  }
  req.Header.Add("Authorization", secret)
  resp, err := client.Do(req)
  if err != nil {
    log.Errorf(ctx, "Failed to fetch API data")
    return make([]byte, 0)
  }
  defer resp.Body.Close()
  news, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Errorf(ctx, "Failed to read API message")
    return make([]byte, 0)
  }
  return news
}
