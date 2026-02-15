package main

import (
	"fmt"
	"time"
)

func main() {
	request, err := NewHttpRequestBuilder("https://api.example.com/users").
		SetMethod(POST).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{"name": "Juan"}).
		SetTimeout(3 * time.Second).
		SetRetries(2).
		Build()

	if err != nil {
		panic(err)
	}

	var response map[string]any
	err = request.Send(&response)
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
