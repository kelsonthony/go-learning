package main

import (
	"errors"
	"fmt"
)

func main() {
	//int8, int16, int32, int64
	var numero int64 = 100000000
	var numero2 int = 100000000
	//var numero3 := 100000000 // tipo inferido

	//uint8, uint16, uint32, uint64
	var numero4 uint32 = 100000000

	//fmt.Println("Numero 3", numero3)

	numero5 := -1000 // tipo inferido

	//float32, float64
	var numeroReal1 float32 = 10.50
	var numeroReal2 float64 = 1000000000.99
	numeroReal3 := 150.99 // tipo inferido

	//uintptr é um tipo especial usado para armazenar endereços de memória
	uint := uintptr(123456)

	// alias de tipos
	var nmero6 rune = 123456 // exatamaente igual ao int32

	// byte é um alias para uint8, usado para representar dados binários
	var numero7 byte = 123

	var str string = "Texto"

	str2 := "Outro texto" // tipo inferido

	// o mais perto do char é o byte ou rune
	char := 'B' // tipo inferido, na verdade é um rune e vai devolvar um valor da tabela ASCII

	// valores zero
	var texto string

	var booleano1 bool = true

	var booleano2 bool // valor zero é false

	var erro error // valor zero é nil

	var err2 error = errors.New("Erro interno")

	fmt.Println("O valor do erro 2 é:", err2)

	fmt.Println("O valor do erro é:", erro)

	fmt.Println("O valor do booleano1 é:", booleano1)
	fmt.Println("O valor do booleano2 é:", booleano2) // valor zero de um boolean é false
	fmt.Println("O valor do texto é:", texto)         // valor zero de uma string é vazia
	fmt.Println("O valor do numero é:", numero)
	fmt.Println("O valor do char é:", char)
	fmt.Println("O valor da string é:", str)
	fmt.Println("O valor da string 2 é:", str2)
	fmt.Println("O valor do numero é:", numero)
	fmt.Println("O valor do numero2 é:", numero2)
	fmt.Println("O valor do numero4 é:", numero4)
	fmt.Println("O valor do numero5 é:", numero5)
	fmt.Println("O valor do numeroReal1 é:", numeroReal1)
	fmt.Println("O valor do numeroReal2 é:", numeroReal2)
	fmt.Println("O valor do numeroReal3 é:", numeroReal3)
	fmt.Println("O valor do uintptr é:", uint)
	fmt.Println("O valor do rune é:", nmero6)
	fmt.Println("O valor do byte é:", numero7)

}
