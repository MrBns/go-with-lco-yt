package main

import (
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	httpClient := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	res, err := httpClient.Do(request)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(data))
}
