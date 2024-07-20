package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func diaDaSemana2(numero int) string {
	var dia string
	switch {
	case numero == 1:
		dia = "Domingo"
	case numero == 2:
		dia = "Segunda"
	case numero == 3:
		dia = "Terça"
	case numero == 4:
		dia = "Quarta"
	case numero == 5:
		dia = "Quinta"
	case numero == 6:
		dia = "Sexta"
	case numero == 7:
		dia = "Sábado"
	default:
		dia = "Número inválido"
	}
	return dia
}

func main() {
	fmt.Println(diaDaSemana(1))  // Output: Domingo
	fmt.Println(diaDaSemana(4))  // Output: Quarta
	fmt.Println(diaDaSemana(10)) // Output: Número inválido

	fmt.Println(diaDaSemana2(1)) // Output: Domingo
}
