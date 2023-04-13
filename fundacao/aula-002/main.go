package main

var (
	saldo       float64 = 123.45
	limite      int     = 120
	nomeUsuario string  = "João"
	ativo       bool    = true
)

func main() {

	println(
		"Variaveis com escopo Globais:",
		"\n saldo: ", saldo,
		"\n limite: ", limite,
		"\n nomeUsuario: ", nomeUsuario,
		"\n ativo: ", ativo,
		"\n",
	)

	// const nao mudam de valor
	println("Constantes: ")
	const msg = "Hello, World!"
	println(msg, "\n")

	// var pode mudar de valor
	println("Variaveis: ")
	var nome string = "João"
	var idade = 18
	var altura = 1.80
	var casado = false
	println(
		"Nome: ", nome,
		"\n idade: ", idade,
		"\n altura: ", altura,
		"\n casado: ", casado,
		"\n",
	)

	nome = "Maria"
	idade = 20
	altura = 1.60
	casado = true
	println(
		"Nome: ", nome,
		"\n idade: ", idade,
		"\n altura: ", altura,
		"\n casado: ", casado,
		"\n",
	)

	// variaval nao usada
	var naoUsada = "Nao Usada"
	println(
		"variaval naoUsada, usada para nao dar erro: ", naoUsada,
		"\n",
	)
	// variaval naoUsada, usada para nao dar erro

	// shortHand
	shortHand := "Short Hand"
	println(
		"Defina uma variavel com shortHand: temAgua := false ", shortHand,
		"\n",
	)

}
