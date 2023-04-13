package main

// Funcoes

func main() {

	println("somar(1, 2) = ", somar(1, 2))
	println("subtrair(1, 2) = ", subtrair(1, 2))
	println("multiplicar(1, 2) = ", multiplicar(1, 2))
	println("dividir(1, 2) = ", dividir(1, 2))
	println("resto(1, 2) = ", resto(1, 2))
	println("retornoMultiplasValores() = ")
	println(retornoMultiplasValores())

}

func somar(a int, b int) int {
	return a + b
}

func subtrair(a int, b int) int {
	return a - b
}

func multiplicar(a int, b int) int {
	return a * b
}

func dividir(a int, b int) int {
	return a / b
}

func resto(a int, b int) int {
	return a % b
}

func retornoMultiplasValores() (int, int) {
	return 1, 2
}
