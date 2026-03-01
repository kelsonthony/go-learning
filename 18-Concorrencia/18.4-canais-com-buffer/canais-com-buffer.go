package main

import "fmt"

// a definição de um canal com buffer é feita passando o tamanho do buffer
// como segundo argumento para a função make.
// O buffer permite que o canal armazene um número específico de mensagens sem bloquear a goroutine
// que está enviando as mensagens. No exemplo abaixo, o canal tem um buffer de 2,
// o que significa que ele pode armazenar até 2 mensagens antes de bloquear a goroutine que está enviando as mensagens.
func main() {
	canal := make(chan string, 2) //canal com buffer de 2
	canal <- "Olá Mundo"
	canal <- "Outra mensagem"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)

}
