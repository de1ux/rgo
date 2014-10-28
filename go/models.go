package server

import (
	//"appengine/datastore"
	"time"
)

type Board struct {
	Name   string   `datastore:"name"`
	Topics *[]Topic `datastore:"-"`
}

type RequiredBoardParams struct {
	Name string
}

type Topic struct {
	Title       string `datastore:"title"`
	Posts       []*Post
	Author      string
	DateCreated time.Time
}

type Post struct {
	Author      string
	Body        []byte
	DateCreated time.Time
	Parent      *Post
}

type Account struct {
	ID          int
	Fullname    string
	Username    string
	DateCreated time.Time
}
