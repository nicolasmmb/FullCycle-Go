package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", BuscarCEPHandler)
	http.ListenAndServe(":8080", nil)

}

func BuscarCEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "O parâmetro 'cep' é obrigatório"}`))
		return
	}
}
