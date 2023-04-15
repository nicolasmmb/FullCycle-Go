package main

import (
	"encoding/json"
	"os"
)

// as tags json são usadas para que o json seja decodificado corretamente no struct
// e também para que o json seja codificado corretamente conforme o nome das tags
type Conta struct {
	Numero int     `json:"numero"`
	Saldo  float64 `json:"saldo"`
}

func main() {
	conta := Conta{Numero: 1406, Saldo: 1000.25}

	// Atribuindo a uma variável
	resultado, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	println(string(resultado))

	// Criando um encoder que escreve no os.Stdout
	encoder := json.NewEncoder(os.Stdout)

	// Escrevendo o objeto conta no encoder
	err = encoder.Encode(conta)

	// Verificando se houve algum erro
	if err != nil {
		panic(err)
	}

	// Lendo Json Puro
	// Deve usar as tags json para que o json seja decodificado corretamente no struct
	// caso contrário, coso não haja tags, o json será decodificado para o nome do campo do struct
	// se houver erros no tags ou campo da struct, o json não será decodificado

	jsonPuro := []byte(`{"numero": 1406, "saldo": 1000.25}`)

	// Criando um decoder que lê do jsonPuro
	var contaDecodificada Conta

	// Decodificando o jsonPuro no objeto contaDecodificada
	// recebe um ponteiro para o objeto que receberá o json decodificado
	err = json.Unmarshal(jsonPuro, &contaDecodificada)

	// Verificando se houve algum erro
	if err != nil {
		panic(err)
	}

	// Imprimindo o objeto contaDecodificada
	println(
		"Conta Decodificada: \n",
		"Numero:", contaDecodificada.Numero,
		" - ",
		"Saldo:", contaDecodificada.Saldo,
	)

}
