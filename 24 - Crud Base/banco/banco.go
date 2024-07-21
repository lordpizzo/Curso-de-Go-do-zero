package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o mysql
)

func Conectar() (*sql.DB, error) {
	urlConexao := "root:spfc1987@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", urlConexao)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	// Configurações adicionais do banco

	return db, nil
}
