package main

import "fmt"

func main() {

	//aritméticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	//atribuição

	var variavel string = "String"
	variavel2 := "String2"

	fmt.Println(variavel, variavel2)

	//operadores relacionais
	maior := 1 > 2
	maiorIgual := 1 >= 2
	igual := 1 == 2
	maiorOuIgual := 1 <= 2
	igual2 := 1 != 2
	menor := 1 < 2

	fmt.Println(maior, maiorIgual, igual, maiorOuIgual, igual2, menor)

	//operadores lógicos
	valorEnd := maior && igual
	ou := maior || igual
	nao := !maior

	fmt.Println(valorEnd, ou, nao)

	//unários
	contador := 0
	contador++
	contador += 2
	contador--
	contador -= 2
	contador *= 2
	contador /= 2
	contador %= 2
	contador <<= 2
	contador >>= 2
	contador &= 2
	contador |= 2
	contador ^= 2
	contador = -contador
	contador = +contador
	contador = ^contador

	//Ternário não tem
}
