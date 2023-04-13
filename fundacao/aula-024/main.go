package main

import "fmt"

// Condicionais
func main() {

	fmt.Println("Condicionais")

	a := 1
	b := 2
	c := 3

	if a > b {
		fmt.Println("a é maior que b")
	}
	if a > c {
		fmt.Println("a é maior que c")
	}
	if b > c {
		fmt.Println("b é maior que c")
	}

	if a > b && a > c {
		fmt.Println("a é maior que b e c")
	}

	if a > b || a > c {
		fmt.Println("a é maior que b ou c")
	}

	switch a {
	case 1:
		fmt.Println("a é igual a 1")
	case 2:
		fmt.Println("a é igual a 2")
	case 3:
		fmt.Println("a é igual a 3")
	default:
		fmt.Println("a é igual a ", a)
	}

}
