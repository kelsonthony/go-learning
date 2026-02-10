package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main!")
	auxiliar.Escrever()
	err := checkmail.ValidateFormat("kelsonthony@gmail.com")
	fmt.Println(err)
}
