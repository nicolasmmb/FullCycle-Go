package main

import "fmt"

// Percorrendo Arrays

func main() {

	// array de inteiros com quantidade de elementos definida
	tamanhoDefinido := [5]int{1, 2, 3, 4, 5}
	fmt.Println(tamanhoDefinido)

	// array de inteiros com quantidade de elementos indefinida
	tamanhoIndefinido := []int{1, 2, 3, 4, 5}
	fmt.Println(tamanhoIndefinido)

	// percorrendo array com for
	for index, value := range tamanhoDefinido {
		fmt.Println("index: ", index, "value: ", value)
	}

}
