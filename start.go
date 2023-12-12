package main

import (
"net/http"
"github.com/go-chi/chi/v5"
"github.com/go-chi/chi/v5/middleware"
)



func main() {

userpwd :=  map[string]string{
"olfox": "A78GHKJG#",
"olfox2": "tuxpux7",    }
	
r := chi.NewRouter()
r.Use(middleware.BasicAuth("url-shortener", userpwd ))
	
r.Get("/", func(w http.ResponseWriter, r *http.Request) {
w.Write([]byte("Welcome"))
})


http.ListenAndServe(":3002", r)
											}
