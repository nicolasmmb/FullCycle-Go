package main

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

// Trabalhando com Post

func main() {
	c := http.Client{
		Timeout: time.Second * 5,
	}
	dadosJson := bytes.NewBuffer([]byte(`{ "nome": "Asuka Langley Soryu", "idade": 14 }`))

	resp, err := c.Post("https://httpbin.org/post", "application/json", dadosJson)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))
}
