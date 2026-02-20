package main

import (
	"fmt"
	"time"
)

func main() {
	//CONCORRÊNCIA != PARALELISMO
	go escreve("Olá Mundo") //goroutine principal
	escreve("Outro texto")
}

func escreve(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
