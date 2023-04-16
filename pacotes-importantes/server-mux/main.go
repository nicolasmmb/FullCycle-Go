package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/home", HomeHandler)
	mux.Handle("/blog", &Blog{Title: "Meu Blog"})

	http.ListenAndServe("0.0.0.0:8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}

type Blog struct {
	Title string
}

func (b *Blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.Title))
}
