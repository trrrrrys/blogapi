package datastore

import "cloud.google.com/go/datastore"

const (
	kindUser    = "User"
	kindContent = "Content"
)

type contentEntity struct {
	Title string
	// Author        *datastore.Key
	Tags          []string
	PublishedDate int
	Body          string
}

var ck = datastore.NameKey("Content", "", nil)
