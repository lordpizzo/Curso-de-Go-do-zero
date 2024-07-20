package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"João", 20}
	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1)
	fmt.Println(e1.nome)
	fmt.Println(e1.idade)
	fmt.Println(e1.curso)
	fmt.Println(e1.faculdade)
}
