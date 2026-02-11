package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arquivo arrays")

	var array1 [5]int

	fmt.Println(array1)

	array2 := [5]int{1, 2, 3, 4, 5}

	fmt.Println(array2)

	// outra forma de criar um array
	array3 := [...]int{1, 2, 3, 4, 5}

	fmt.Println(array3)

	// outra forma de criar um array
	array4 := [5]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}

	fmt.Println(array4)

	var array5 [5]string

	array5[0] = "A"
	array5[1] = "B"
	array5[2] = "C"
	array5[3] = "D"
	array5[4] = "E"
	fmt.Println(array5)

	array6 := [5]string{"A", "B", "C", "D", "E"}
	fmt.Println(array6)

	array7 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array7)

	fmt.Println(len(array7))

	// array7[5] = 6 // erro: index out of range [5] with length 5
	//array6[5] = "F"

	array8 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array8)

	//array8[5] = 6 // erro: cannot assign to array8[5] (index out of range [5] with length 5)

	//slices são mais flexíveis que arrays, pois eles podem crescer e diminuir de tamanho dinamicamente, enquanto os arrays têm um tamanho fixo
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(slice)

	slice = append(slice, 6)

	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array8))

	slice2 := array6[1:3]
	fmt.Println(slice2)

	slice2[1] = "X"
	fmt.Println(slice2)
	fmt.Println(array6)

}
