package main

import "fmt"

// Maps

func main() {

	nomeIdade := map[string]int{}
	nomeIdade["Juan"] = 1
	nomeIdade["Nicolas"] = 23
	nomeIdade["Maria"] = 45

	// printando nome e idade
	println("Printando nome e idade")
	for nome, idade := range nomeIdade {
		fmt.Printf("Nome: %s \t - \tIdade: %d \t\n", nome, idade)
	}

	// duplicando idade
	println("Duplicando idade")
	for nome, idade := range nomeIdade {
		nomeIdade[nome] = idade * 2
	}

	for nome, idade := range nomeIdade {
		fmt.Printf("Nome: %s \t - \tIdade: %d \t\n", nome, idade)
	}
}
