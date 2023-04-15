package main

import (
	"io"
	"net/http"
)

func main() {
	// Realizando uma chamada HTTP
	request, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	// Lendo o corpo da resposta
	result, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	// Imprimindo o resultado
	println(string(result))

	// Fechando o corpo da resposta
	request.Body.Close()
}
