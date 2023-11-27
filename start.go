package main
# d12344
import (
"net/http"
"github.com/go-chi/chi/v5"
"github.com/go-chi/chi/v5/middleware")

func main() {
r := chi.NewRouter()
r.Use(middleware.Logger)
r.Get("/", func(w http.ResponseWriter, r *http.Request) {
w.Write([]byte("welcome"))
})
http.ListenAndServe(":3000", r)
											}
