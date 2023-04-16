package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCEP struct {
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	CEP         string `json:"cep"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

func main() {

	http.HandleFunc("/", BuscarCEPHandler)
	http.ListenAndServe(":8080", nil)

}

func BuscarCEPHandler(w http.ResponseWriter, r *http.Request) {
	// Verificando se a URL é a raiz
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// Definindo o tipo de conteúdo da resposta
	w.Header().Set("Content-Type", "application/json")

	// Pegando o parâmetro 'cep' da URL
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "O parâmetro 'cep' é obrigatório"}`))
		return
	}

	// Buscando o CEP
	cep, error := BuscarCEP(cepParam)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Ocorreu um erro ao buscar o CEP"}`))
		return
	}

	// Escrevendo a resposta
	w.WriteHeader(http.StatusOK)

	// Escrevendo o JSON
	json.NewEncoder(w).Encode(cep)

}

func BuscarCEP(cep string) (*ViaCEP, error) {
	resp, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()

	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}

	var c ViaCEP
	error = json.Unmarshal(body, &c)
	if error != nil {
		return nil, error
	}

	return &c, nil
}
