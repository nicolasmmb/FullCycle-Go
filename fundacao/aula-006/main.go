package main

import "fmt"

// Slices

func main() {

	itens := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Todos: ", itens)
	fmt.Println("Dois primeiros: ", itens[:2])
	fmt.Println("Dois últimos: ", itens[8:])

	// adicionando itens no slice
	novosItens := append(itens, 11, 12, 13, 14, 15)
	fmt.Println("Novo: ", novosItens)
	fmt.Println("Antigos: ", itens)

}
