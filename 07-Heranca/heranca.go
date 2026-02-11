package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

// heranca: a struct estudante tem acesso aos campos da struct pessoa, ou seja, ela herda os campos da struct pessoa
// pessoa declarada dentro de estudante, ou seja, estudante tem uma pessoa dentro dela, ou seja, estudante é uma pessoa, ou seja, estudante herda os campos da struct pessoa
// não precisa espeficar o tipo no ex: pessoa pessoa, basta declarar pessoa, ou seja, estudante tem acesso aos campos da struct pessoa, ou seja, estudante herda os campos da struct pessoa
type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Arquivo structs")

	p1 := pessoa{nome: "João", sobrenome: "Silva", idade: 30, altura: 180}
	fmt.Printf("O nome da pessoa é %s %s, ela tem %d anos e mede %d cm\n", p1.nome, p1.sobrenome, p1.idade, p1.altura)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Printf("O nome do estudante é %s %s, ele tem %d anos, mede %d cm, estuda %s na faculdade %s\n", e1.nome, e1.sobrenome, e1.idade, e1.altura, e1.curso, e1.faculdade)

	// outra forma de criar um estudante
	e2 := estudante{pessoa: pessoa{nome: "Pedro", sobrenome: "Almeida", idade: 20, altura: 170}, curso: "Medicina", faculdade: "Unicamp"}
	fmt.Printf("O nome do estudante é %s %s, ele tem %d anos, mede %d cm, estuda %s na faculdade %s\n", e2.nome, e2.sobrenome, e2.idade, e2.altura, e2.curso, e2.faculdade)
}
