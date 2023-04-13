package main

// Ponteiros

func main() {

	// variavel -> endereço de memória -> valor
	var numero int = 10
	var ptrNumero *int = &numero

	println("Endereço de memória de numero: ", &numero)
	println("Endereço de memória de ptrNumero: ", ptrNumero)

	println("Valor de numero: ", numero)
	println("Valor de ptrNumero: ", *ptrNumero)

	println("Alterando o valor de ptrNumero que é o endereço de memória de 'numero'")
	*ptrNumero = 20
	println("Valor de numero: ", numero)
	println("Valor de ptrNumero: ", *ptrNumero)

	println("O 'numero' e o 'ptrNumero' apontam para o mesmo endereço de memória")
	println("Se alterar o valor de 'numero' o 'ptrNumero' também será alterado")
	println("Se alterar o valor de 'ptrNumero' o 'numero' também será alterado")

}
