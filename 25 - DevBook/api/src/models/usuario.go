package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Apelido  string    `json:"apelido,omitempty"`
	Email    string    `json:"email,omitempty"`
	Telefone string    `json:"telefone,omitempty"`
	Admin    bool      `json:"admin,omitempty"`
	Ativo    bool      `json:"ativo,omitempty"`
	Avatar   string    `json:"avatar,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("O Email é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("Email inválido")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("O Senha é obrigatório e não pode estar em branco")
	}

	if usuario.Apelido == "" {
		return errors.New("O Apelido é obrigatório e não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Apelido = strings.TrimSpace(usuario.Apelido)
	usuario.Telefone = strings.TrimSpace(usuario.Telefone)

	if etapa == "cadastro" {
		senhaComHash, erro := security.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}
		usuario.Senha = string(senhaComHash)

	}
	return nil
}
