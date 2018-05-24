package main

import (
  "encoding/json"
  "net/http"
  "time"
  "google.golang.org/appengine"
  "google.golang.org/appengine/log"
  "github.com/unrolled/render"
)

// Ping status for Go server
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Ping string }{"OK"})
	}
}

// ETL top news stories
func etlHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var response Response
		response.Status = "FAILURE"
		ctx := appengine.NewContext(req)
		newsPayload := fetchNews(ctx)
		if len(newsPayload) > 0 {
			var news News
			err := json.Unmarshal(newsPayload, &news)
			if err != nil {
				log.Errorf(ctx, "Failed to deserialize news content <%v>", err)
			} else {
				response.Volume = len(news.Articles)
				ctn := buildContent(news.Articles)
				if save(ctx, ctn); err != nil {
					log.Errorf(ctx, "Failed to persist content <%v>", err)
				} else {
					response.Status = "SUCCESS"
				}
			}
		} else {
			log.Errorf(ctx, "Failed to retrieve news content")
		}
		formatter.JSON(w, http.StatusOK, response)
	}
}

// Remove expired articles
func destroyHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := appengine.NewContext(req)
		s := http.StatusNoContent
		t := time.Now()
		cutoff := t.AddDate(0, 0, -7)
		if err := delete(ctx, cutoff); err != nil {
          log.Errorf(ctx, "Failed to clean the data cache")
          s = http.StatusInternalServerError
		}
		formatter.JSON(w, s, nil)
	}
}