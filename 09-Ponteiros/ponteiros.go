package main

import "fmt"

func main() {
	fmt.Println("Arquivo ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// Ponteiro é uma referência de memória
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)

	//desreferenciando o ponteiro
	fmt.Println(variavel3, *ponteiro)

	//modificando o valor da variável3 através do ponteiro
	variavel3 = 150
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro)

}
