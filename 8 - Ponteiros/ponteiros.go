package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	// & é usado para pegar o endereço de memória
	var variavel2 int = variavel1
	// * é usado para pegar o valor da variável
	fmt.Println(variavel1, variavel2)
	// & é usado para pegar o endereço de memória
	var variavel3 *int = &variavel1
	// * é usado para pegar o valor da variável
	fmt.Println(variavel3)  // mostra o endereço
	fmt.Println(*variavel3) // mostra o valor
	// & é usado para pegar o endereço de memória

}
