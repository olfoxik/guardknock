package main

import (
"net/http"
"github.com/go-chi/chi/v5"
"github.com/go-chi/chi/v5/middleware"
"gopkg.in/ini.v1"
"fmt"
"os"
)



func main() {

cfg, err := ini.Load("config.toml")
if err != nil {
fmt.Printf("Fail to read file: %v", err)
os.Exit(1)     }

userpwd :=  map[string]string{
"olfox": "A78GHKJG#",
"olfox2": "tuxpux7",    }
	
r := chi.NewRouter()
r.Use(middleware.BasicAuth("url-shortener", userpwd ))
	
r.Get("/", func(w http.ResponseWriter, r *http.Request) {
w.Write([]byte("111") )
})

 

	
http.ListenAndServe(":3002", r)
											}
