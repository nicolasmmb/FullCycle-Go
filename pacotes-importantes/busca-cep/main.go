package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
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
	file, err := os.OpenFile("./cep.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro na criação do arquivo: %v \n", err)
	}
	defer file.Close()

	for _, cep := range os.Args[1:] {

		cep = "https://viacep.com.br/ws/" + cep + `/json`
		request, err := http.Get(cep)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro na requisição: %v \n", err)
		}
		defer request.Body.Close()

		response, err := io.ReadAll(request.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro na leitura da resposta: %v \n", err)
		}

		fmt.Printf("%s \n", response)

		var cep ViaCEP

		err = json.Unmarshal(response, &cep)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro na decodificação do json: %v \n", err)
		}

		fmt.Printf("CEP: %s, Logradouro: %s, Bairro: %s, Localidade: %s, UF: %s \n",
			cep.CEP,
			cep.Logradouro,
			cep.Bairro,
			cep.Localidade,
			cep.UF,
		)

		// add line to csv
		_, err = file.WriteString(fmt.Sprintf("%s,%s,%s,%s,%s \n", cep.CEP, cep.Logradouro, cep.Bairro, cep.Localidade, cep.UF))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro na escrita do arquivo: %v \n", err)
		}

	}
	time.Sleep(10 * time.Second)
	os.Remove("./cep.csv")
}
