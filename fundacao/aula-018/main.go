package main

import "fmt"

// Interfaces vazias

func main() {
	//Evitar ao maximo usar interfaces vazias

	var x interface{} = 1
	var y interface{} = "Nicolas"
	var z interface{} = true

	printInterface(x)
	printInterface(y)
	printInterface(z)
}

func printInterface(x interface{}) {
	// tipo e valores
	fmt.Printf("Tipo: %T \t Valor: %v\n", x, x)
}
