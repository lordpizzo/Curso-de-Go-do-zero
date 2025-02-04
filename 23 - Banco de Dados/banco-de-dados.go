package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	urlConexao := "root:spfc1987@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", urlConexao)
	//defer db.Close()
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()
	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}
	//fmt.Println("Conexão aberta!")
	println(db)

	linhas, erro := db.Query("select * from usuarios")

	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()
	fmt.Println(linhas)
}
