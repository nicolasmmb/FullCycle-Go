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
	// defer: executa a função no final do escopo
	// defer é chamado quando a função termina
	defer request.Body.Close()

	// Lendo o corpo da resposta
	result, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	// Imprimindo o resultado
	println(string(result))

	// Fechando o corpo da resposta
	request.Body.Close()

	// exemplo de defer
	deferExample()
}

func deferExample() {
	defer println("End")
	println("Start")
	println("Middle")

}
