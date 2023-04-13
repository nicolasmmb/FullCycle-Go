package main

// Generics

func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

// Constrains para generics
type Numero interface {
	int | float64
}

func SomaGenerico[T Numero](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	mInteiro := map[string]int{
		"Juan":    1,
		"Nicolas": 23,
		"Maria":   45,
	}
	mFloat := map[string]float64{
		"Juan":    1.1,
		"Nicolas": 23.23,
		"Maria":   45.45,
	}
	println(SomaInteiro(mInteiro))
	println(SomaFloat(mFloat))

	println("Somando exemplo de generics")
	println(SomaGenerico(mInteiro))
	println(SomaGenerico(mFloat))

}
