package main

import "fmt"

func main() {
	canal := make(chan string, 2) //canal com buffer de 2
	canal <- "Olá Mundo"
	canal <- "Outra mensagem"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)

}
