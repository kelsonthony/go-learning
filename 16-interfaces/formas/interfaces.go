package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	largura, altura float64
}

type circulo struct {
	raio float64
}

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("Área da forma: %.2f\n", f.area())
}

func (r retangulo) area() float64 {
	return r.largura * r.altura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	r := retangulo{largura: 5, altura: 3}
	escreverArea(r)

	c := circulo{raio: 10}
	escreverArea(c)
}
