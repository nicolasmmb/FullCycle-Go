package main

// For
func main() {
	// tipos de for

	// for comum
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(numeros); i++ {
		println(numeros[i])
	}

	// for com range
	for i, numero := range numeros {
		println(i, numero)
	}

	// for com blank identifier
	for _, numero := range numeros {
		println(numero)
	}

	// for while
	i := 10
	for i > 0 {
		println(i)
		i--
	}

	// for infinito
	for {
		println("loop infinito")
	}

}
