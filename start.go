package main

import (
"net/http"
"github.com/go-chi/chi/v5"
"github.com/go-chi/chi/v5/middleware")

func main() {
r := chi.NewRouter()
r.Use(middleware.BasicAuth("url-shortener", map[string]string{
  "olfox": "mirumir57#",
	"olfox2": "tuxpux7",
     }))
r.Get("/", func(w http.ResponseWriter, r *http.Request) {
w.Write([]byte("welcome"))
})


http.ListenAndServe(":3002", r)
											}
