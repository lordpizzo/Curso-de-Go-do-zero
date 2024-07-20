package main

import (
	"errors"
	"fmt"
)

func main() {
	//todos os tipos de inteiros possíveis para GO int8, int16, int32, int64 --- int utiliza o tipo referente a arquitetura do computador
	var numero int16 = 100
	var numero2 int = -100
	fmt.Println(numero)
	fmt.Println(numero2)

	//uint é um inteiro positivo
	var numero3 uint = 100
	fmt.Println(numero3)

	//alias
	//INT32 = RUNE
	var numero4 rune = 12345
	fmt.Println(numero4)

	//alias BYTE = UINT8
	var numero5 byte = 123
	fmt.Println(numero5)

	//float
	//float32, float64
	var numero6 float32 = 123.45

	fmt.Println(numero6)

	//Strings
	var str string = "texto"
	fmt.Println(str)

	//Char
	var char rune = 'A'
	fmt.Println(char)

	// FIM STRING - String vazia, Numeros iguais a 0

	var texto string
	fmt.Println(texto) //vazio

	//boolean
	var booleano1 bool = true
	fmt.Println(booleano1)

	//ERROR - Valor zero é nil
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
