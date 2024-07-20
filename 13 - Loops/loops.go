package main

import (
	"fmt"
)

func main() {
	i := 0

	for i < 10 {
		// time.Sleep(time.Second)
		i++
		fmt.Println(i)
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
		// time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}
	// for i := 0; i < len(nomes); i++ {
	// 	fmt.Println(nomes[i])
	// }

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "Silva",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for chave := range usuario {
		fmt.Println(chave, usuario[chave])
	}

}
