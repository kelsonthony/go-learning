package enderecos

// Teste de Unidade
import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoEndereco(t *testing.T) {

	t.Parallel()

	cenarioDeTestes := []cenarioDeTeste{
		{enderecoInserido: "Avenida Paulista", retornoEsperado: "Avenida"},
		{enderecoInserido: "Estrada do Sol", retornoEsperado: "Tipo de endereço inválido"},
		{enderecoInserido: " ", retornoEsperado: "Tipo de endereço inválido"},
		{enderecoInserido: "Rua das Flores", retornoEsperado: "Rua"},
		{enderecoInserido: "alameda dos Anjos", retornoEsperado: "Alameda"},
	}

	for _, cenario := range cenarioDeTestes {
		tipoEndereco := TipoEndereco(cenario.enderecoInserido)

		if tipoEndereco != cenario.retornoEsperado {
			t.Errorf("O tipo do endereço '%s' deveria ser '%s', mas foi '%s'", cenario.enderecoInserido, cenario.retornoEsperado, tipoEndereco)
		}
	}
}

func TestTipoEnderecoComRua(t *testing.T) {

	t.Parallel()

	tipoEndereco := TipoEndereco("Avenida Paulista")

	if tipoEndereco != "Avenida" {
		t.Errorf("O tipo do endereço deveria ser 'Avenida', mas foi '%s'", tipoEndereco)
	}
}

func TestTipoEnderecoComTipoInvalido(t *testing.T) {

	t.Parallel()

	tipoEndereco := TipoEndereco("Estrada do Sol")

	if tipoEndereco != "Tipo de endereço inválido" {
		t.Errorf("O tipo do endereço deveria ser 'Tipo de endereço inválido', mas foi '%s'", tipoEndereco)
	}
}

func TestTipoEnderecoComLetraMaiuscula(t *testing.T) {

	t.Parallel()

	tipoEndereco := TipoEndereco("Rua das Flores")

	if tipoEndereco != "Rua" {
		t.Errorf("O tipo do endereço deveria ser 'Rua', mas foi '%s'", tipoEndereco)
	}
}

func TestTipoEnderecoComLetraMinuscula(t *testing.T) {

	t.Parallel()

	tipoEndereco := TipoEndereco("alameda dos Anjos")

	if tipoEndereco != "Alameda" {
		t.Errorf("O tipo do endereço deveria ser 'Alameda', mas foi '%s'", tipoEndereco)
	}
}
