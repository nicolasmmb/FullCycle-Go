package main

import (
	"fmt"

	"github.com/google/uuid"
)

// Instalando pacotes

// Passos para baixar um pacote externo
// opcional:
// 	-> go mod init github.com/nicolasmmb/go-expert
// go get github.com/google/uuid
// go mod tidy

func main() {
	id := uuid.New()
	fmt.Println(id)
}
