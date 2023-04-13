package main

// Criação de tipos

type Pessoa struct {
	nome   string
	idade  int
	altura float64
	casado bool
}

func main() {

	pessoa := Pessoa{
		nome:   "Nicolas",
		idade:  23,
		altura: 1.78,
		casado: false,
	}

	println(
		"\n nome: ", pessoa.nome,
		"\n idade: ", pessoa.idade,
		"\n altura: ", pessoa.altura,
		"\n casado: ", pessoa.casado,
		"\n",
	)
}
