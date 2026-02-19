package main

import "fmt"

func funcao1() {
	fmt.Println("Função 1")
}

func funcao2() {
	fmt.Println("Função 2")
}

func alunoEstaAprovado(n1, n2 float64) bool {
	// defer é muito útil para manipular conexões com banco de dados, arquivos, etc.
	// Ele garante que a função adiada seja executada mesmo que ocorra um erro ou uma condição de saída antecipada na função principal.
	defer fmt.Println("Média calculada. Resultado será retornado!")
	fmt.Println("Calculando a média...")

	media := (n1 + n2) / 2

	if media >= 6 {

		return true
	} else {

		return false
	}
}

func main() {
	fmt.Println("Função Defer")
	//defer é utilizado para adiar a execução de uma função até que a função que a chamou termine
	defer funcao1()
	funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))
}
