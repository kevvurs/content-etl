package main

import "time"

type News struct {
  Status        string     `json:"status"`
  TotalResults  int64      `json:"totalResults"`
  Articles      []Article  `json:"articles"`
}

type Article struct {
  Source       Source    `json:"source"`
  Author       string    `json:"author"`
  Title        string    `json:"title"`
  Description  string    `json:"description"`
  Url          string    `json:"url"`
  UrlToImage   string    `json:"urlToImage"`
  PublishedAt  string    `json:"publishedAt"`
}

type Source struct {
  Id    string  `json:"id"`
  Name  string  `json:"name"`
}

type Response struct {
  Volume  int     `json:"volume"`
  Status  string  `json:"status"`
}

type Content struct {
  Id           string     `datastore:"id"`
  Source       string     `datastore:"source"`
  Author       string     `datastore:"author"`
  Title        string     `datastore:"title"`
  Description  string     `datastore:"description,noindex"`
  Url          string     `datastore:"url"`
  UrlToImage   string     `datastore:"urlToImage"`
  PublishedAt  time.Time  `datastore:"publishedAt"`
}
