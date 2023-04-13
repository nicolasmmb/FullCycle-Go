package main

import "fmt"

// Funções variádicas

func main() {
	fmt.Println("aceita qualquer quantidade de argumentos definidos no tipo da função variádica")
	fmt.Println(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// exemplo de função variádica
	fmt.Println("sumList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) --> ", sumList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

}

func sumList(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}
