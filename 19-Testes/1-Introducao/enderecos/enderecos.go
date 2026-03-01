package enderecos

import "strings"

// TipoEndereco verifica se um endereço tem um tipo válido e o retorna. Se o tipo for inválido, retorna uma mensagem de erro.
func TipoEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "alameda", "travessa", "praça"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false

	for _, tipoValido := range tiposValidos {
		if tipoValido == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
		if enderecoTemUmTipoValido {
			return strings.Title(primeiraPalavraDoEndereco)
		}
	}

	return "Tipo de endereço inválido"

}
