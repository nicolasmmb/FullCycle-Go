package main

// Composição de Structs

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Telefone   string
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco Endereco
}

func main() {

	println("Criando a Struct Cliente \n")
	cliente := Cliente{
		Nome:  "Nicolas",
		Idade: 23,
		Ativo: true,
		Endereco: Endereco{
			Logradouro: "Rua 1",
			Numero:     1,
			Cidade:     "São Paulo",
			Telefone:   "11 99999-9999",
		},
	}

	println("Printando cliente")
	println(
		"\n nome: ", cliente.Nome,
		"\n idade: ", cliente.Idade,
		"\n ativo: ", cliente.Ativo,
		"\n logradouro: ", cliente.Endereco.Logradouro,
		"\n numero: ", cliente.Endereco.Numero,
		"\n cidade: ", cliente.Endereco.Cidade,
		"\n telefone: ", cliente.Endereco.Telefone,
		"\n",
	)

}
