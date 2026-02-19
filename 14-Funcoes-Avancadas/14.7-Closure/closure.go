package main

import "fmt"

func closureFunc() func() {
	texto := "Olá, func closure!"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao

}

func main() {

	texto := "Dentro da função main"
	fmt.Println("Sempre vai executar: ", texto)

	funcaoNova := closureFunc()
	funcaoNova()

}
