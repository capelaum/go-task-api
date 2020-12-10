package controller

import (
	"net/http"
)

// Register a multiplexer
func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping())
	mux.HandleFunc("/create", create())
	mux.HandleFunc("/list", list())
	mux.HandleFunc("/list/", listByName())
	mux.HandleFunc("/delete/", delete())
	return mux
}
