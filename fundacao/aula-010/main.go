package main

// Closures

func main() {
	// Closures
	total := func() int {
		return sumList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}
	// exemplo de closure
	println("Exemplo de Closures: total() -->", total())

}

func sumList(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}
