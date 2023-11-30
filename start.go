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
    Database struct {
        Server        string
        Ports         []int
        ConnectionMax uint
        Enabled       bool
    }
    Servers map[string]ServerInfo
    Clients struct {
        Data  [][]interface{}
        Hosts []string
    }
}

type ServerInfo struct {
    IP net.IP
    DC string
}

func main() {

  f, err := os.Open("config.toml")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    var config tomlConfig
    if err := toml.NewDecoder(f).Decode(&config); err != nil {
        panic(err)
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
