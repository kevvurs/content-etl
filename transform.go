package main

import (
  "crypto/md5"
  "fmt"
  "time"
  "golang.org/x/net/context"
)

func buildContent(ctx context.Context, arts []Article) []*Content {
  content := make([]*Content, len(arts))
  for idx, art := range arts {
    content[idx] = polish(&art)
  }
  return content
}

func polish(art *Article) *Content {
  con := new(Content)
  con.Id = hashId(art.Source.Name, art.Author, art.Title)
  con.Source = art.Source.Name
  con.Author = art.Author
  con.Title = art.Title
  con.Description = art.Description
  con.Url = art.Url
  con.UrlToImage = art.UrlToImage
  con.PublishedAt = parseTimestamp(art.PublishedAt)
  return con
}

func hashId(source, author, title string) string {
  contentId := fmt.Sprintf("%s:%s:%s", source, author, title)
  h := md5.New()
  b := h.Sum([]byte(contentId))
  return string(b)
}

func parseTimestamp(ts string) time.Time {
  t, _ := time.Parse(time.RFC3339, ts)
  return t
}
