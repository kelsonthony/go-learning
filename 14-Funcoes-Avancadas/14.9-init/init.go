package main

import "fmt"

var n int

// A função init é uma função especial em Go que é executada automaticamente antes da função main.
// Ela é usada para inicializar variáveis, configurar o ambiente ou realizar qualquer configuração necessária antes do programa começar
// a executar a lógica principal. V
// Você pode ter várias funções init em um pacote, e elas serão executadas na ordem em que aparecem no código.
// Pode ser útil para configurar conexões de banco de dados, inicializar variáveis globais ou realizar qualquer
// configuração necessária antes do programa começar a executar a lógica principal.
// dentro de cada arquivo .go, a função init será executada antes da função main, e se houver várias funções init em diferentes arquivos,
// elas serão executadas na ordem em que os arquivos são compilados.
func init() {
	fmt.Println("Hello, Func init!")
	n = 10
}

func main() {
	fmt.Println("Hello, Func main!")
	fmt.Println(n)
}
