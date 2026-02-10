package main

// toda variavel que for declarada deve ser utilizada
func main() {
	var nome string = "Kelson"
	nome2 := "Mchilanny" // tipo inferido
	var idade int = 39
	var altura float32 = 1.88
	var estudante bool = true

	var (
		var1 string = "valor1"
		var2 int    = 2
		var3 bool   = false
	)

	var4, var5, var6 := "valor4", 5, true

	const constante string = "Eu sou uma constante"

	var1, var4 = var4, var1 // trocando valores de var5 e var6

	println("var1 recebe var4:", var1)
	println("var4 recebe var1:", var4)

	println(constante)

	println("Var1:", var1)
	println("Var2:", var2)
	println("Var3:", var3)
	println("Var4:", var4)
	println("Var5:", var5)
	println("Var6:", var6)

	println("Nome:", nome)
	println("Idade:", idade)
	println("Altura:", altura)
	println("Estudante:", estudante)
	println("Nome2:", nome2)
}
