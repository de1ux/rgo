package server

import (
	"time"
)

type Thread struct {
	Title       string `datastore:"title"`
	Posts       []Post
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
