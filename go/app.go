package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/api/board/{resourceId}", BoardHandler)
	r.HandleFunc("/api/boards/", BoardListHandler)

	r.HandleFunc("/api/thread/{resourceId}", ThreadHandler)
	r.HandleFunc("/api/threads/", ThreadListHandler)

	r.HandleFunc("/api/post/{resourceId}", PostHandler)
	r.HandleFunc("/api/posts/", PostListHandler)

	http.Handle("/", r)
}
