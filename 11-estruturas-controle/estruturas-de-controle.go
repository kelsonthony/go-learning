package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 10

	// Estrutura de controle if-else
	// Não é necessário usar parênteses em torno da condição, mas as chaves são obrigatórias
	if numero > 15 {
		fmt.Println("O número é maior que 15")
	} else {
		fmt.Println("O número é menor ou igual a 15")
	}

	//if init; condition { ... }
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("O número é maior que zero")
	} else if numero < -10 {
		fmt.Println("O número é menor que -10")
	} else {
		fmt.Println("O número é zero")
	}

	// A variável outroNumero só existe dentro do bloco if-else, então não pode ser acessada aqui fora
	//fmt.Println(outroNumero)
}
