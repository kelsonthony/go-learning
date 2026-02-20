package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		escreve("Olá Mundo")
		waitGroup.Done()
	}()

	go func() {
		escreve("Outro texto")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escreve(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
