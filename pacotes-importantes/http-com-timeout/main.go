package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

// HttpClient com Timeout

func main() {
	c := http.Client{
		Timeout: time.Second * 2,
	}

	resp, err := c.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))
}
