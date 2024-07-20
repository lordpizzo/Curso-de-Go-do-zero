package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	fmt.Println(u)

	u.nome = "João"
	u.idade = 30
	u.endereco = endereco{"Rua dos Bobos", 0}
	//u := usuario{"João", 30, "Rua dos Bobos, 0"}
	fmt.Println(u)

	u2 := usuario{nome: "Maria", idade: 25, endereco: endereco{"Rua dos Bobos", 0}}
	fmt.Println(u2)

	u3 := usuario{nome: "José"}
	fmt.Println(u3)
	fmt.Println(u3.nome)
	fmt.Println(u3.idade) // erro, idade não é um campo definido
}
