package main

import "fmt"

func recuperarExecucao() {
	fmt.Println("Tentativa de recuperar a execução")
	//o valor nill indica que não houve um panic, ou seja, a execução ocorreu normalmente
	//nill é null em outras linguagens de programação
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	fmt.Println("Calculando a média...")
	media := (n1 + n2) / 2

	if media > 6 {
		fmt.Println("Aluno Aprovado")
		return true
	} else if media < 6 {
		fmt.Println("Aluno Reprovado")
		return false
	}

	//o panic é utilizado para indicar que ocorreu um erro grave e que a execução do programa deve ser interrompida.
	// Ele é comumente usado para lidar com situações inesperadas ou condições de erro que não podem ser tratadas de forma adequada.
	panic("Média é exatamente 6.")

}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pos execução")

}
