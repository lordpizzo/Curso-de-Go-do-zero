package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario cria um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		w.Write(([]byte("Falha ao ler o corpo da requisição!")))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	fmt.Println(usuario)

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")

	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)

	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	idInserido, erro := insercao.LastInsertId()

	if erro != nil {
		w.Write([]byte("Erro ao obter o ID inserido"))
		return
	}
	//status code
	w.WriteHeader(http.StatusCreated)
	_, erro = w.Write([]byte(fmt.Sprintf("Criando usuário no servidor! id: %d", idInserido)))
}

// BuscarUsuarios busca todos os usuarios no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")

	if erro != nil {
		w.Write([]byte("Erro ao buscar usuários"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao ler os dados do usuário"))
			return
		}
		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuarios)
}

// BuscarUsuario busca um detemrinado usuario a partir de seu id
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	id, erro := strconv.ParseUint(parametros["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro ID para inteiro"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	linha, erro := db.Query("select * from usuarios where id = ?", id)

	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuário"))
		return
	}
	defer linha.Close()

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao ler os dados do usuário"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuario)
}

// AtualizarUsuario realiza a atualização de um usuario
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	id, erro := strconv.ParseUint(parametros["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro ID para inteiro"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		w.Write(([]byte("Falha ao ler o corpo da requisição!")))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")

	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	defer statement.Close()

	fmt.Println(" O que tem no usuario")
	fmt.Println(usuario)

	if _, erro = statement.Exec(usuario.Nome, usuario.Email, id); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeletarUsuario remove um usuario do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	id, erro := strconv.ParseUint(parametros["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro ID para inteiro"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")

	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	defer statement.Close()

	if _, erro = statement.Exec(id); erro != nil {
		w.Write([]byte("Erro ao deletar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
