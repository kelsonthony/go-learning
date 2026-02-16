package main

func calculosMatematicos(a, b int) (soma, subtracao, multiplicacao, divisao int) {
	soma = a + b
	subtracao = a - b
	multiplicacao = a * b
	divisao = a / b

	return
}

func main() {
	somaResultado, subtracaoResultado, multiplicacaoResultado, divisaoResultado := calculosMatematicos(10, 5)

	println("Soma:", somaResultado)
	println("Subtração:", subtracaoResultado)
	println("Multiplicação:", multiplicacaoResultado)
	println("Divisão:", divisaoResultado)

}
