package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Teste de área do retângulo", func(t *testing.T) {
		r := Retangulo{largura: 5, altura: 3}
		areaRetangulo := r.Area()
		if areaRetangulo != 15 {
			t.Fatalf("Área do retângulo incorreta. Esperado: 15, Obtido: %.2f", areaRetangulo)
		}
	})

	t.Run("Teste de área do círculo", func(t *testing.T) {
		c := Circulo{raio: 10}
		areaCirculo := c.Area()
		expectedAreaCirculo := math.Pi * math.Pow(10, 2)
		if areaCirculo != expectedAreaCirculo {
			t.Fatalf("Área do círculo incorreta. Esperado: %.2f, Obtido: %.2f", expectedAreaCirculo, areaCirculo)
		}
	})

}
