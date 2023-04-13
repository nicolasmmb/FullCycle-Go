package main

// Iniciando com Structs

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	println("Criando a Struct Cliente \n")
	cliente := Cliente{
		Nome:  "Nicolas",
		Idade: 23,
		Ativo: true,
	}

	println("Printando cliente")
	println(
		"\n nome: ", cliente.Nome,
		"\n idade: ", cliente.Idade,
		"\n ativo: ", cliente.Ativo,
		"\n",
	)

}
