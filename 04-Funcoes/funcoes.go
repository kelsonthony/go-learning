package main

import "fmt"

func main() {
	soma := soma(10, 20)

	var f = func(txt string) string {
		fmt.Println("Função F")
		return txt
	}

	println(soma)

	resultado := f("Texto da função anônima")
	fmt.Println(resultado)

	resultadosSoma, resultadosSubtracao, resultadosMultiplicacao, resultadosDivisao := calculosMatematicos(20, 10)
	fmt.Println(resultadosSoma, resultadosSubtracao, resultadosMultiplicacao, resultadosDivisao)

	// ignorar um ou mais retornos usando o underline
	// apenas o resultado da soma será armazenado, os outros serão ignorados
	//a ordem dos retornos deve ser mantida, ou seja, o resultado da soma deve ser o primeiro retorno
	resultadosSoma2, _, _, _ := calculosMatematicos(20, 10)
	fmt.Println(resultadosSoma2)

	_, _, resultadosMultiplicacao2, _ := calculosMatematicos(20, 10)
	fmt.Println(resultadosMultiplicacao2)

}

// o tipo apos os parametros é o tipo de retorno da função
func soma(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// funcao com multiplos retornos, os tipos de retorno devem ser nomeados
func calculosMatematicos(n1, n2 int64) (soma, subtracao, multiplicacao, divisao int64) {
	soma = n1 + n2
	subtracao = n1 - n2
	multiplicacao = n1 * n2
	divisao = n1 / n2

	return soma, subtracao, multiplicacao, divisao
}
