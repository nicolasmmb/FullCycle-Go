package main

import (
	"bufio"
	"fmt"
	"os"
)

const path = "pacotes-importantes/manipulacao-de-arquivos/arquivo.txt"

func main() {
	// Abrindo um arquivo
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	// Escrevendo no arquivo
	tamanho, err := f.WriteString("Escrevendo no arquivo")
	if err != nil {
		panic(err)
	}

	// Exibindo o nome do arquivo e o tamanho
	fmt.Printf("Arquivo aberto com sucesso: %v, com tamanho de %d bytes \n", f.Name(), tamanho)

	// Lendo o arquivo
	arquivo, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	// Exibindo o conteúdo do arquivo
	fmt.Println(string(arquivo))

	// Fechando o arquivo
	err = f.Close()
	if err != nil {
		panic(err)
	}

	// Lendo o arquivo otimizado para memoria
	arquivoLeitura, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	// Lendo o arquivo com buffer
	reader := bufio.NewReader(arquivoLeitura)

	// Define o tamanho do buffer(Quantidade de bytes que serão lidos)
	buffer := make([]byte, 11)

	// Lendo o arquivo
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		// Exibindo o conteúdo do arquivo
		fmt.Println(string(buffer[:n]))
	}

	// Fechando o arquivo
	err = arquivoLeitura.Close()
	if err != nil {
		panic(err)
	}

	// removendo o arquivo
	err = os.Remove(path)

}
