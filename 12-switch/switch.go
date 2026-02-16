package main

import "fmt"

func diaDaSemana(dia int) string {
	switch dia {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func diaDaSemana2(dia int) string {
	var diadaSemana string
	switch {
	case dia == 1:
		diadaSemana = "Domingo"
		//fallthrough // Força a execução do próximo case, mesmo que a condição seja falsa
		//break // Interrompe a execução do switch, mesmo que haja mais cases para serem avaliados
	case dia == 2:
		diadaSemana = "Segunda-feira"
	case dia == 3:
		diadaSemana = "Terça-feira"
	case dia == 4:
		diadaSemana = "Quarta-feira"
	case dia == 5:
		diadaSemana = "Quinta-feira"
	case dia == 6:
		diadaSemana = "Sexta-feira"
	case dia == 7:
		diadaSemana = "Sábado"
	default:
		diadaSemana = "Número inválido"
	}
	return diadaSemana
}

func main() {
	fmt.Println("Estrutura de controle switch")
	diaAtual := diaDaSemana(6)
	fmt.Println("Hoje é o dia: " + diaAtual)

	fmt.Println("Usando switch de outra forma")
	diaAtual2 := diaDaSemana2(1)
	fmt.Println("Hoje é o dia: " + diaAtual2)
}
