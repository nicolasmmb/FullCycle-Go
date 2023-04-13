package main

// Type assertation

func main() {

	var nome interface{} = "Nicolas"
	var idade interface{} = 23
	// sem type assertation
	println("Sem type assertation: ", nome)
	// type assert
	println("Com type assertation: ", nome.(string))

	// sem type assertation
	println("Sem type assertation: ", idade)
	// type assert
	println("Com type assertation: ", idade.(int))

}
