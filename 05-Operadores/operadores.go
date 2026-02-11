package main

import "fmt"

func main() {
	// Operadores aritméticos
	soma := 10 + 20
	subtracao := 20 - 10
	multiplicacao := 10 * 20
	divisao := 20 / 10
	resto := 20 % 3

	fmt.Println("Soma:", soma)
	fmt.Println("Subtração:", subtracao)
	fmt.Println("Multiplicação:", multiplicacao)
	fmt.Println("Divisão:", divisao)
	fmt.Println("Resto:", resto)

	// Operadores de atribuição
	var x int = 10
	x += 5 // x = x + 5
	fmt.Println("Valor de x após += :", x)

	//operadores de atribuição simples
	x = 10
	x -= 3 // x = x - 3
	fmt.Println("Valor de x após -= :", x)

	x = 10
	x *= 2 // x = x * 2
	fmt.Println("Valor de x após *= :", x)

	x = 10
	x /= 2 // x = x / 2
	fmt.Println("Valor de x após /= :", x)

	x = 10
	x %= 3 // x = x % 3
	fmt.Println("Valor de x após %= :", x)

	// fim dos operadores de relacionais, os outros são -=, *=, /=, %=

	// Operadores de comparação
	fmt.Println("10 == 20:", 10 == 20)
	fmt.Println("10 != 20:", 10 != 20)
	fmt.Println("10 > 20:", 10 > 20)
	fmt.Println("10 < 20:", 10 < 20)
	fmt.Println("10 >= 20:", 10 >= 20)
	fmt.Println("10 <= 20:", 10 <= 20)
	// fim dos operadores de relacionais

	//operadores lógicos
	fmt.Println("true && false:", true && false)
	fmt.Println("true || false:", true || false)
	fmt.Println("!true:", !true)
	// fim dos operadores lógicos

	//operadores de incremento e decremento - unarios
	var y int = 10
	y++ // y = y + 1
	fmt.Println("Valor de y após incremento:", y)

	y-- // y = y - 1
	fmt.Println("Valor de y após decremento:", y)
	// fim dos operadores de incremento e decremento

	y *= 2 // y = y * 2
	fmt.Println("Valor de y após multiplicação:", y)

	y /= 2 // y = y / 2
	fmt.Println("Valor de y após divisão:", y)

	y %= 3 // y = y % 3
	fmt.Println("Valor de y após resto:", y)
	// fim dos operadores de atribuição simples

	// em Go não existe operador de incremento ou decremento prefixado, apenas o posfixado
	// ou seja, não existe ++y ou --y, apenas y++ e y--

	// Em go não existe operador ternário, ou seja, não existe a sintaxe condicao ? valor1 : valor2
	// para simular um operador ternário, podemos usar uma estrutura condicional if-else
	idade := 18
	var resultado string
	if idade >= 18 {
		resultado = "Maior de idade"
	} else {
		resultado = "Menor de idade"
	}
	fmt.Println(resultado)

}
