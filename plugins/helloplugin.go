package main

import (
	"C"
	"net/http"
)

func ServeHTTP (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HelloPlugin"))
}