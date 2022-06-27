package main

import (
	"fmt"

	"github.com/sanketrai1/vault/secret"
)

func main() {
	vault, err := secret.Vault("password")
	if err != nil {
		panic(err)
	}
	cipher, err := vault.Set("twitter", "my_key")
	if err != nil {
		panic(err)
	}
	fmt.Println(cipher)
	value, err := vault.Get("twitter")
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
}
