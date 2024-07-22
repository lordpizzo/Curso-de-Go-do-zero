package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePuclicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao models.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (models.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
		SELECT p.id, p.titulo, p.conteudo, p.autor_id, u.apelido
		FROM publicacoes p
		INNER JOIN usuarios u ON u.id = p.autor_id
		WHERE p.id = ?
	`, publicacaoID)
	if erro != nil {
		return models.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao models.Publicacao

	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.AutorApelido,
		); erro != nil {
			return models.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

func (respositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]models.Publicacao, error) {
	fmt.Println(fmt.Sprintf("Zent : %d", usuarioID))
	linhas, erro := respositorio.db.Query(`
		SELECT DISTINCt p.*, u.apelido from publicacoes p
		inner join usuarios u on u.id = p.autor_id
		inner join seguidores s on p.autor_id = s.seguidor_id
			where u.id = ? or s.seguidor_id = ?
			order by 1 desc
	`, usuarioID, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {
		var publicacao models.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorApelido,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repositorio Publicacoes) BuscarPorIdUsuario(usuarioID uint64) ([]models.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT p.*, u.apelido FROM publicacoes p
		JOIN usuarios u ON u.id = p.autor_id
		WHERE p.autor_id = ?
	`, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {
		var publicacao models.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorApelido,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao models.Publicacao) error {
	statement, erro := repositorio.db.Prepare(
		"UPDATE publicacoes SET titulo = ?, conteudo = ? WHERE id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID)
	if erro != nil {
		return erro
	}

	return nil
}

func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"DELETE FROM publicacoes WHERE id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(publicacaoID)
	if erro != nil {
		return erro
	}

	return nil
}

func (repositorio Publicacoes) Curtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"UPDATE publicacoes SET curtidas = curtidas + 1 WHERE id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(publicacaoID)
	if erro != nil {
		return erro
	}

	return nil
}

func (repositorio Publicacoes) Descurtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"UPDATE publicacoes SET curtidas = curtidas - 1 WHERE id = ? AND curtidas > 0",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(publicacaoID)
	if erro != nil {
		return erro
	}

	return nil
}
