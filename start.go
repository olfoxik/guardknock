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
hh := cfg.Section("users").KeysHash()
	
if err != nil {
fmt.Printf("Fail to read file: %v", err)
os.Exit(1)     }

//userpwd :=  map[string]string{
//"olfox": "A78GHKJG#",
//"olfox2": "tuxpux7",    }
fmt.Printf (hh)
r := chi.NewRouter()
r.Use(middleware.BasicAuth("url-shortener", hh ))
	
r.Get("/", func(w http.ResponseWriter, r *http.Request) {
w.Write([]byte("www") )

})

 

	
http.ListenAndServe(":3002", r)
											}
