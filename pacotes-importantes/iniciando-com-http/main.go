package main

import "net/http"

func main() {

	http.HandleFunc("/buscar-cep", BuscarCEP)
	http.ListenAndServe(":8080", nil)
}

func BuscarCEP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando CEP..."))
}
