package enderecos

import "strings"

// TipoDeEndereco verifica se um endereço tem um tipo válido
func TipoDeEndereco(enderecos string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoEmLetraMinuscula := strings.ToLower(enderecos)
	primeiraPalavra := strings.Split(enderecoEmLetraMinuscula, " ")[0]
	for _, tipo := range tiposValidos {
		if primeiraPalavra == tipo {
			return "Tipo válido"
		}
	}
	return "Tipo inválido"
}
