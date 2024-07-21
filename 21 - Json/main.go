package main

import (
	"encoding/json"
	"fmt"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJson := `{"nome":"Rex","raca":"vira-lata","idade":3}`
	var c cachorro = cachorro{"Rex", "vira-lata", 3}
	json.Unmarshal([]byte(cachorroEmJson), &c)
	fmt.Println(c)

	cachorroEmJson2, erro := json.Marshal(c)
	if erro != nil {
		fmt.Println(erro)
	}
	fmt.Println(string(cachorroEmJson2))

	var c2 cachorro
	fmt.Println(c2)
	json.Unmarshal([]byte(cachorroEmJson), &c2)
	//c2 = c
	fmt.Println(c2)
}
