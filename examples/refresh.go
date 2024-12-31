package main

import (
	"fmt"
	"github.com/beclab/lldap-client/pkg/auth"
)

func main() {
	refreshToken := "d2g5LSMXjRZJJa7qhWVftjhMfuTepqL2unNGvkBwPvKf8vw1c5EEzNvELQtReBICOm25jCefh78kx12AmxALNaW7VjOZhlLO3p9W+test001"
	t, err := auth.Refresh("http://127.0.0.1:17170", refreshToken)
	if err != nil {
		panic(err)
	}
	fmt.Println("token: ", t)
}
