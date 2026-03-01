package formas

import (
	"fmt"
	"math"
)

type Retangulo struct {
	largura, altura float64
}

type Circulo struct {
	raio float64
}

type Forma interface {
	Area() float64
}

func EscreverArea(f Forma) {
	fmt.Printf("Área da forma: %.2f\n", f.Area())
}

func (r Retangulo) Area() float64 {
	return r.largura * r.altura
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	r := Retangulo{largura: 5, altura: 3}
	EscreverArea(r)

	c := Circulo{raio: 10}
	EscreverArea(c)
}
