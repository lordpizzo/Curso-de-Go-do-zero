package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)

}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {

	usuario1 := usuario{"João", 25}

	println(usuario1.nome)

	usuario1.nome = "Pedro"
	println(usuario1.nome)

	usuario3 := usuario{"João", 25}
	println(usuario3.nome)
	println(usuario3.idade)

	usuario1.salvar()

	fmt.Println(usuario1.maiorDeIdade())
	fmt.Println(usuario1)
	usuario1.fazerAniversario()

	fmt.Println(usuario1)
}
