package main

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


router.Route("/url", func(r chi.Router) {
    // Подключаем авторизацию
    r.Use(middleware.BasicAuth("url-shortener", map[string]string{
        // Передаем в middleware креды
        "eee": "vvvv",
        // Если у вас более одного пользователя,
        // то можете добавить остальные пары по аналогии.
    }))

    r.Post("/", save.New(log, storage))
})

http.ListenAndServe(":3000", r)
											}
