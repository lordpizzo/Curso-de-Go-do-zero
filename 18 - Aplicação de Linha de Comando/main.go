package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("ponto de partida")

	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args)

	if erro != nil {
		log.Fatal(erro)
	}
}
