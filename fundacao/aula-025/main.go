package main

import "fmt"

// Compilando o projeto

func main() {

	fmt.Println("Compilando o projeto")
	fmt.Println("go build")
	fmt.Println("go build -o nome_do_executavel")
	fmt.Println("go tool dist list")
	fmt.Println("GOOS=linux GOARCH=amd64 go build -o nome_do_executavel")
	fmt.Println("GOOS=windows GOARCH=amd64 go build -o nome_do_executavel.exe")
	fmt.Println("GOOS=darwin GOARCH=amd64 go build -o nome_do_executavel")

}
