package main

// Interfaces

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

type Pessoa interface {
	desativar()
}

func (c Cliente) imprimir() {
	println(
		" nome: ", c.Nome,
		"\n idade: ", c.Idade,
		"\n ativo: ", c.Ativo,
		"\n telefone: ", c.Endereco.Telefone,
		"\n",
	)
}

func (c *Cliente) desativar() {
	println("Desativando o cliente: ")
	c.Ativo = false
}

func Desativacao(p Pessoa) {
	p.desativar()
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
	cliente.imprimir()

	Desativacao(&cliente)
	cliente.imprimir()

}
