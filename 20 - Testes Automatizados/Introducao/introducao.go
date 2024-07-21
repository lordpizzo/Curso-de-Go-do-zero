package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoDeEndereco := enderecos.TipoDeEndereco("Rodovia dos imigrantes")
	fmt.Println(tipoDeEndereco)
	fmt.Println("Executado")
}
