package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// Usuarios representa o repositório de usuários
type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuário no banco de dados
func (repositorio Usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, apelido, email, senha, telefone, avatar, admin, ativo) values (?, ?, ?, ?, ?, ?, ?, ?)",
	)

	if erro != nil {
		return 0, erro
	}

	resultado, erro := statement.Exec(usuario.Nome, usuario.Apelido, usuario.Email, usuario.Senha, usuario.Telefone, usuario.Avatar, usuario.Admin, usuario.Ativo)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

// Buscar busca um usuário baseado numa query param
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, erro := repositorio.db.Query(
		"select id, nome, apelido, email, telefone, avatar, admin, ativo from usuarios where nome like ? or apelido like ?",
		nomeOuNick, nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []models.Usuario
	for linhas.Next() {
		var usuario models.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Apelido,
			&usuario.Email,
			&usuario.Telefone,
			&usuario.Avatar,
			&usuario.Admin,
			&usuario.Ativo,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarPorID busca um usuário Baseado no id
func (repositorio Usuarios) BuscarPorID(ID uint64) (models.Usuario, error) {
	linha, erro := repositorio.db.Query(
		"select id, nome, apelido, email, telefone, avatar, admin, ativo from usuarios where id = ?",
		ID,
	)

	if erro != nil {
		return models.Usuario{}, erro
	}

	defer linha.Close()

	var usuario models.Usuario
	if linha.Next() {
		if erro := linha.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Apelido,
			&usuario.Email,
			&usuario.Telefone,
			&usuario.Avatar,
			&usuario.Admin,
			&usuario.Ativo,
		); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Atualizar atualiza um usuário
func (repositorio Usuarios) Atualizar(ID uint64, usuario models.Usuario) error {
	statement, erro := repositorio.db.Prepare(
		"update usuarios set nome = ?, apelido = ?, email = ?, telefone = ?, avatar = ?, admin = ?, ativo = ? where id = ?",
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Apelido, usuario.Email, usuario.Telefone, usuario.Avatar, usuario.Admin, usuario.Ativo, ID); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui um usuário pelo ID
func (repositorio Usuarios) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"delete from usuarios where id = ?",
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	linha, erro := repositorio.db.Query(
		"select id, senha from usuarios where email = ?",
		email,
	)

	if erro != nil {
		return models.Usuario{}, erro
	}

	defer linha.Close()

	var usuario models.Usuario
	if linha.Next() {
		if erro := linha.Scan(
			&usuario.ID,
			&usuario.Senha,
		); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)",
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"delete from seguidores where usuario_id = ? and seguidor_id = ?",
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) BuscarSeguidores(usuarioID uint64) ([]models.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select u.id, u.nome, u.apelido, u.email, u.telefone, u.avatar, u.admin, u.ativo "+
			"from usuarios u inner join seguidores s on u.id = s.seguidor_id where s.usuario_id = ?",
		usuarioID,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []models.Usuario
	for linhas.Next() {
		var usuario models.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Apelido,
			&usuario.Email,
			&usuario.Telefone,
			&usuario.Avatar,
			&usuario.Admin,
			&usuario.Ativo,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (respositorio Usuarios) BuscarSeguindo(usuarioID uint64) ([]models.Usuario, error) {
	linhas, erro := respositorio.db.Query(
		"select u.id, u.nome, u.apelido, u.email, u.telefone, u.avatar, u.admin, u.ativo "+
			"from usuarios u inner join seguidores s on u.id = s.usuario_id where s.seguidor_id = ?",
		usuarioID,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []models.Usuario
	for linhas.Next() {
		var usuario models.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Apelido,
			&usuario.Email,
			&usuario.Telefone,
			&usuario.Avatar,
			&usuario.Admin,
			&usuario.Ativo,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repoositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repoositorio.db.Query(
		"select senha from usuarios where id = ?",
		usuarioID,
	)

	if erro != nil {
		return "", erro
	}

	defer linha.Close()

	var usuario models.Usuario
	if linha.Next() {
		if erro := linha.Scan(
			&usuario.Senha,
		); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil
}

func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statement, erro := repositorio.db.Prepare(
		"update usuarios set senha = ? where id = ?",
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(senha, usuarioID); erro != nil {
		return erro
	}

	return nil
}
