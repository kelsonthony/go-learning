package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	fmt.Printf("A soma dos números é: %d\n", total)
	return total
}

func escrever(texto string, numeros ...int) {

	for _, numero := range numeros {
		fmt.Printf("%s %d\n", texto, numero)
	}
}

func main() {
	soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	//posso não passar nenhum número, e a função ainda funcionará e retornará 0
	//soma()
	escrever("Olá mundo", 1, 2, 4, 4)
}
