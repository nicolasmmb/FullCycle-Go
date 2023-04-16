package main

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

// Customizando uma requisição

func main() {
	c := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(
		"POST",
		"https://httpbin.org/post",
		bytes.NewBufferString("Nome=João"),
	)
	if err != nil {
		panic(err)
	}

	req.Header.Set("X-Accept", "application/json")
	req.Header.Set("X-Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-Connection", "close")
	// é possível modificar varios headers e outras propriedades da requisição

	resp, err := c.Do(req)
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
