package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string) //criando um canal

	go escrever("Olá Mundo", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	// for {
	// 	mensagem, aberto := <-canal //recebendo dados do canal
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do Programa")

}

func escrever(texto string, canal chan string) {
	//time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- texto //enviando dados para o canal
		time.Sleep(time.Second)
	}
	close(canal)
}
