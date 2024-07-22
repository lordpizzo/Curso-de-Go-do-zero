package database

import (
	"api/src/config"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		log.Fatal(erro)
		return nil, erro
	}

	if erro := db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
