package main

import "fmt"

func main() {
	retorno := func(texto string) string {
		return fmt.Sprintf("O texto é: %s", texto)
	}("Olá Mundo")

	fmt.Println(retorno)
}
