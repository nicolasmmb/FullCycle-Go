package main

// Ponteiros e Structs

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func (c *Cliente) ativarComPonteiro() {
	println("Ativando o cliente com ponteiro: ")
	c.Ativo = true
}

func (c Cliente) ativarSemPonteiro() {
	println("Ativando o cliente sem ponteiro: ")
	c.Ativo = true
}

func PrintCliente(c Cliente) {
	println(
		" nome: ", c.Nome,
		"\n idade: ", c.Idade,
		"\n ativo: ", c.Ativo,
		"\n",
	)
}

func main() {

	println("Criando a Struct Cliente \n")
	cliente := Cliente{
		Nome:  "Nicolas",
		Idade: 23,
		Ativo: true,
	}

	println("Imprimindo o cliente: \n")
	cliente.Ativo = false
	PrintCliente(cliente)

	cliente.Ativo = false
	cliente.ativarComPonteiro()
	PrintCliente(cliente)

	cliente.Ativo = false
	cliente.ativarSemPonteiro()
	PrintCliente(cliente)

}
