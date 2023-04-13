package main

import "fmt"

// Importando fmt e tipagem

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

	fmt.Printf(
		"\n nome: %s \n idade: %d \n altura: %f \n casado: %t \n",
		pessoa.nome,
		pessoa.idade,
		pessoa.altura,
		pessoa.casado,
	)

	// print dos tipos
	fmt.Printf(
		"\n nome: %T \n idade: %T \n altura: %T \n casado: %T \n",
		pessoa.nome,
		pessoa.idade,
		pessoa.altura,
		pessoa.casado,
	)
}
