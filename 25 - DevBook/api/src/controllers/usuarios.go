package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("cadastro"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usuarioID, erro := repositorio.Criar(usuario)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	usuario.ID = usuarioID
	response.JSON(w, http.StatusCreated, usuario)
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))
	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usuarios, erro := repositorio.Buscar(nomeOuNick)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, usuarios)
}

func BuscarUsuarioPorID(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)

	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usuario, erro := repositorio.BuscarPorID(usuarioID)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, usuario)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDNoToken, erro := auth.ExtrairUsuarioID(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	if usuarioID != usuarioIDNoToken {
		response.Erro(w, http.StatusForbidden, errors.New("Ação invalida. Você não term essa permissão"))
		return
	}
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("edicao"); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	if erro := repositorio.Atualizar(usuarioID, usuario); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDNoToken, erro := auth.ExtrairUsuarioID(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	if usuarioID != usuarioIDNoToken {
		response.Erro(w, http.StatusForbidden, errors.New("Ação invalida. Você não term essa permissão"))
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Deletar(usuarioID); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorID, erro := auth.ExtrairUsuarioID(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorID == usuarioID {
		response.Erro(w, http.StatusForbidden, errors.New("Você não pode seguir a si mesmo"))
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Seguir(usuarioID, seguidorID); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorID, erro := auth.ExtrairUsuarioID(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorID == usuarioID {
		response.Erro(w, http.StatusForbidden, errors.New("Você não pode parar de seguir a si mesmo"))
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.PararDeSeguir(usuarioID, seguidorID); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	seguidores, erro := repositorio.BuscarSeguidores(usuarioID)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, seguidores)
}

func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	seguindo, erro := repositorio.BuscarSeguindo(usuarioID)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, seguindo)
}

func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	usuarioIDNoToken, erro := auth.ExtrairUsuarioID(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if usuarioIDNoToken != usuarioID {
		response.Erro(w, http.StatusForbidden, errors.New("Você não pode atualizar a senha de outro usuário"))
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var senha models.Senha
	if erro = json.Unmarshal(corpoRequisicao, &senha); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	senhaSalvaNoBanco, erro := repositorio.BuscarSenha(usuarioID)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.VerificarSenha(senhaSalvaNoBanco, senha.Atual); erro != nil {
		response.Erro(w, http.StatusUnauthorized, errors.New("Senha inválida"))
		return
	}

	senhaComHash, erro := security.Hash(senha.Nova)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.AtualizarSenha(usuarioID, string(senhaComHash)); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
