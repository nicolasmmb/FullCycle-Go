package main

// Quando usar ponteiros

func somar(
	a int,
	b int,
) int {
	return a + b
}

func somarComPonteirosMaisCinco(
	a *int,
	b *int,
) int {
	*a = *a + 5
	return *a + *b
}

func main() {
	variaval01 := 1
	variaval02 := 2

	println("variaval01 = ", variaval01)
	println("variaval02 = ", variaval02)

	println("somar(1, 2) = ", somar(variaval01, variaval02))

	println("somarComPonteirosMaisCinco(1, 2) = ", somarComPonteirosMaisCinco(&variaval01, &variaval02))

	println("Passadno o enderedo da memoria ao soma com ponteiros, os valores das variaveis mudam")
	println("variaval01 = ", variaval01)
	println("variaval02 = ", variaval02)

	println(
		"OBS: Quando passamos o endereço de memória de uma variável para uma função, " +
			"podemos alterar o valor da variável dentro da função" +
			"e o valor da variável fora da função também será alterado",
	)
}
