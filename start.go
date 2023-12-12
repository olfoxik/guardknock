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

cfg, err := ini.Load("config.ini")
userpwd := cfg.Section("users").KeysHash()
fmt.Println("map:", userpwd)	
if err != nil {
fmt.Printf("Fail to read file: %v", err)
os.Exit(1)   }

r := chi.NewRouter()
r.Use(middleware.BasicAuth("url-shortener", userpwd ))
	
r.Get("/", func(w http.ResponseWriter, r *http.Request) {
w.Write([]byte("Welcome GuardKnock v 1.0 "))
})


http.ListenAndServe(":3004", r)
											}
