package main

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"
)

// Trabalhando com http e contextos

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	dadosJson := bytes.NewBuffer([]byte(`{ "nome": "Asuka Langley Soryu", "idade": 14 }`))

	req, err := http.NewRequestWithContext(ctx, "POST", "https://httpbin.org/post", dadosJson)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
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
