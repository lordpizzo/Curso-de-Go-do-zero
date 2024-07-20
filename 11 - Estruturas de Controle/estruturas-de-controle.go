package main

import "fmt"

func main() {
	fmt.Println("estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("maior que 15")
	} else {
		fmt.Println("menor ou igual a 15")
	}

	// if init
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("numero é maior que zero")
	} else {
		fmt.Println("numero é menor ou igual a zero")
	}

	// if e else if
	outraCondicao := false

	if numero > 5 {
		fmt.Println("maior que 5")
	} else if outraCondicao {
		fmt.Println("outra condicao é verdadeira")
	} else {
		fmt.Println("nenhuma das condições anteriores foi atendida")
	}

}
