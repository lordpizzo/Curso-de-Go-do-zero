package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {
	usuarioId, erro := auth.ExtrairUsuarioID(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repositorio := repositories.NovoRepositorioDePuclicacoes(db)
	publicacoes, erro := repositorio.BuscarPorUsuario(usuarioId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusOK, publicacoes)
}

func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
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
	repositorio := repositories.NovoRepositorioDePuclicacoes(db)
	publicacao, erro := repositorio.BuscarPorID(publicacaoID)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusOK, publicacao)

}

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioId, erro := auth.ExtrairUsuarioID(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	var publicacao models.Publicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	publicacao.AutorID = usuarioId

	if erro = publicacao.Preparar(); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	repositorio := repositories.NovoRepositorioDePuclicacoes(db)
	publicacao.ID, erro = repositorio.Criar(publicacao)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusCreated, publicacao)

}

func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioId, erro := auth.ExtrairUsuarioID(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
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
	repositorio := repositories.NovoRepositorioDePuclicacoes(db)
	publicacaoBanco, erro := repositorio.BuscarPorID(publicacaoID)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	if publicacaoBanco.AutorID != usuarioId {
		response.Erro(w, http.StatusForbidden, nil)
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao models.Publicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	publicacao.ID = publicacaoID

	if erro = publicacao.Preparar(); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.Atualizar(publicacao.ID, publicacao); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioId, erro := auth.ExtrairUsuarioID(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
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
	repositorio := repositories.NovoRepositorioDePuclicacoes(db)
	publicacaoBanco, erro := repositorio.BuscarPorID(publicacaoID)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	if publicacaoBanco.AutorID != usuarioId {
		response.Erro(w, http.StatusForbidden, nil)
		return
	}

	if erro = repositorio.Deletar(publicacaoID); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

func BuscarPublicacoesPorUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
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
	repositorio := repositories.NovoRepositorioDePuclicacoes(db)
	publicacoes, erro := repositorio.BuscarPorIdUsuario(usuarioId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusOK, publicacoes)
}

func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
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
	repositorio := repositories.NovoRepositorioDePuclicacoes(db)
	if erro = repositorio.Curtir(publicacaoID); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
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
	repositorio := repositories.NovoRepositorioDePuclicacoes(db)
	if erro = repositorio.Descurtir(publicacaoID); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}
