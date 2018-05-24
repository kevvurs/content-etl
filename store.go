package main

import (
  "time"
  "golang.org/x/net/context"
  "google.golang.org/appengine/datastore"
  "google.golang.org/appengine/log"
)


func save(ctx context.Context, content []*Content) error {
  keys := make([]*datastore.Key, len(content))
  for idx, cn := range content {
    key := datastore.NewKey(ctx, "Article", cn.Id, 0, nil)
    keys[idx] = key
  }
  _, err := datastore.PutMulti(ctx, keys, content)
  return err
}

func delete(ctx context.Context, t time.Time) error {
  var keys []*datastore.Key
  log.Warningf(ctx, "Deleting publications before %s", t.Format(time.RFC3339))
  q := datastore.NewQuery("Article").
    Filter("publishedAt <", t).
    KeysOnly()
  itr := q.Run(ctx)
  for {
    var c Content
    k, err := itr.Next(&c)
    if err == datastore.Done {
      break
    }
    if err != nil {
      return err
    }
    keys = append(keys, k)
  }
  if err := datastore.DeleteMulti(ctx, keys); err != nil {
  	return err
  }
  return nil 
}
