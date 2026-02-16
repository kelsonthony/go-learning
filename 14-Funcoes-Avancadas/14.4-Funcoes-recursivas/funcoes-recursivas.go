package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	// A função chama a si mesma, mas com um valor menor, até chegar na condição de parada
	//Exemplo: fibonacci(4) = fibonacci(2) + fibonacci(3) ou seja, fibonacci(4) = (fibonacci(0) + fibonacci(1)) + (fibonacci(1) + fibonacci(2))
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	// 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
	posicao := uint(12)
	fmt.Printf("O número na posição %d da sequência de Fibonacci é: %d\n", posicao, fibonacci(posicao))

	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
