package main

import (
	"fmt"
)

func main() {
	// for loop
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// // while loop (using for)
	// j := 0
	// for j < 5 {
	// 	fmt.Println(j)
	// 	j++
	// }

	// infinite loop
	// for {
	// 	fmt.Println("This will run forever")
	// }

	// i := 0
	// for i < 10 {
	// 	i++
	// 	time.Sleep(time.Second)
	// 	fmt.Println("Incrementando o i: ", +i)
	// }

	// for k := 0; k < 5; k++ {
	// 	fmt.Println("Valor de k: ", k)
	// 	time.Sleep(time.Second)
	// }

	// for k := 0; k < 5; k += 2 {
	// 	fmt.Println("Valor de k: ", k)
	// 	time.Sleep(time.Second)
	// }

	//range loop
	nomes := []string{"Alice", "Bob", "Charlie"}
	for index, nome := range nomes {
		fmt.Printf("Index: %d, Nome: %s\n", index, nome)
	}

	for _, nome := range nomes {
		fmt.Printf("Somente oNome: %s\n", nome)
	}

	for indice, letra := range "GoLang" {
		fmt.Printf("Índice: %d, Letra: %c\n", indice, letra)
		fmt.Println(indice, letra)         // vai retornar o indice e o valor da letra, mas a letra é retornada como um número (código Unicode)
		fmt.Println(indice, string(letra)) // para converter o número de volta para a letra, usamos string(letra)
	}

	// range loop com map
	idades := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
	}

	for nome, idade := range idades {
		fmt.Printf("Nome: %s, Idade: %d\n", nome, idade)
	}

	// loops não funciona em structs, mas podemos usar range para iterar sobre os campos de um struct usando reflect (avançado)
	// type Pessoa struct {
	// 	Nome  string
	// 	Idade int
	// }

	// pessoa := Pessoa{
	// 	Nome:  "Mchilanny",
	// 	Idade: 44,
	// }

	// usando reflect para iterar sobre os campos de um struct

	// for campo, valor := range pessoa {
	// 	fmt.Printf("Campo: %s, Valor: %v\n", campo, valor)
	// }

}
