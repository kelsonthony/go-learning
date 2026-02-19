package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// o u seria uma variavel para referenciar o usuario que chamou o método, ou seja, o usuario1
func (u usuario) salvar() {
	fmt.Printf("Salvando usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"João", 30}
	usuario1.salvar()

	usuario2 := usuario{"Maria", 25}
	usuario2.salvar()

	maiorDeIdade := usuario1.maiorDeIdade()
	fmt.Printf("O usuário %s é maior de idade? %t\n", usuario1.nome, maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Printf("O usuário %s agora tem %d anos\n", usuario2.nome, usuario2.idade)
}
