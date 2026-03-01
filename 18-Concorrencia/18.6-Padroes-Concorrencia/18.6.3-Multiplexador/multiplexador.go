package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá, mundo!"), escrever("Programando em Go!"))

	for i := 0; i < 10; i++ {
		mensagem := <-canal
		fmt.Println(mensagem)
	}

}

func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalDeEntrada1:
				canalDeSaida <- fmt.Sprintf("Canal 1: %s", mensagem)
			case mensagem := <-canalDeEntrada2:
				canalDeSaida <- fmt.Sprintf("Canal 2: %s", mensagem)
			}
		}
	}()

	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
