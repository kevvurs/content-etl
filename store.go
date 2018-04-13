package main

import (
  "golang.org/x/net/context"
  "google.golang.org/appengine/datastore"
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
