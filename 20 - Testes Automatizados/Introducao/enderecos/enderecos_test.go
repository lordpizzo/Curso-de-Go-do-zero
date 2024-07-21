package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	t.Run("Avenida", func(t *testing.T) {
		enderecoParaTeste := "Avenida Paulista"
		tipoDeEnderecoEsperado := "Avenida"
		tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

		if tipoDeEnderecoRecebido != "Tipo válido" {
			t.Errorf("O tipo recebido é diferente do esperado. Esperava %s e recebi %s", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
		}
	})

	t.Run("Rua", func(t *testing.T) {
		enderecoParaTeste := "Rua dos Bobos"
		tipoDeEnderecoEsperado := "Rua"
		tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

		if tipoDeEnderecoRecebido != "Tipo válido" {
			t.Errorf("O tipo recebido é diferente do esperado. Esperava %s e recebi %s", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
		}
	})
}
