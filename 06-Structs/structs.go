package main

import "fmt"

// equivalente a uma classe em outras linguagens
type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "João"
	u.idade = 30

	u.endereco.logradouro = "Rua A"
	u.endereco.numero = 123

	fmt.Println(u)

	fmt.Printf("O nome do usuário é %s, ele tem %d anos e mora na %s, %d\n", u.nome, u.idade, u.endereco.logradouro, u.endereco.numero)

	// outra forma de criar um struct
	u2 := usuario{nome: "Maria", idade: 25}
	fmt.Printf("O nome do usuário é %s e ele tem %d anos\n", u2.nome, u2.idade)

	// outra forma de criar um struct
	u3 := usuario{"Pedro", 40, endereco{}}
	fmt.Printf("O nome do usuário é %s e ele tem %d anos\n", u3.nome, u3.idade)

	u4 := usuario{idade: 20}
	fmt.Printf("Somente idade do usuário é %d\n", u4.idade)

}
