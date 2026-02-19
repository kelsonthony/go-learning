package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Printf("Número original: %d\n", numero)
	fmt.Printf("Número invertido: %d\n", numeroInvertido)
	fmt.Printf("Número original novamente: %d\n", numero)

	inverterSinalPonteiro(&numero)
	fmt.Printf("Número original após inverter com ponteiro: %d\n", numero)
	fmt.Printf("Número original novamente mas ponteiros: %d\n", numero)
}
