package main

import (
"net/http"
"github.com/go-chi/chi/v5"
"github.com/go-chi/chi/v5/middleware"
"github.com/BurntSushi/toml"
"time"
)

type tomlConfig struct {
    Title string
    Owner struct {
        Name string
        Dob  time.Time
    }

func main() {

var conf Config
if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
        // обработка ошибки.
}

r := chi.NewRouter()
r.Use(middleware.BasicAuth("url-shortener", map[string]string{
"olfox": "A78GHKJG#",
"olfox2": "tuxpux7",
     }))
r.Get("/", func(w http.ResponseWriter, r *http.Request) {
w.Write([]byte("Welcome"))
})


http.ListenAndServe(":3002", r)
											}
