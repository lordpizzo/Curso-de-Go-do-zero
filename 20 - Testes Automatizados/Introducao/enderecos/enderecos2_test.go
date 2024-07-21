package enderecos

import "testing"

type cenarioDeTeste struct {
	endereco string
	retorno  string
}

func TestTipoDeEndereco2(t *testing.T) {

	cenarioDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Tipo válido"},
		{"Avenida Paulista", "Tipo válido"},
		{"Estrada Qualquer", "Tipo válido"},
		{"Rodovia dos Imigrantes", "Tipo válido"},
		{"Praça das Rosas", "Tipo inválido"},
		{"", "Tipo inválido"},
	}

	for _, cenario := range cenarioDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.endereco)
		if tipoDeEnderecoRecebido != cenario.retorno {
			t.Errorf("O tipo recebido é diferente do esperado. Esperava %s e recebi %s", cenario.retorno, tipoDeEnderecoRecebido)
		}
	}

}
