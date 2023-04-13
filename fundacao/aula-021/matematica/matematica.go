package matematica

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
