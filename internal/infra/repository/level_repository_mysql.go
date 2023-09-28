package repository

import (
	"database/sql"

	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type LevelRepositoryMysql struct {
	db *sql.DB
}

func NewLevelRepositoryMysql(db *sql.DB) *LevelRepositoryMysql {
	return &LevelRepositoryMysql{
		db: db,
	}
}

func (cr *LevelRepositoryMysql) FindById(id string) (*entity.Level, error) {
	var level entity.Level

	sqlStatement := `
		SELECT
			id,
			name,
			permissions,
			company_id
		FROM levels
		WHERE id = ?
			AND deleted_at IS NULL
	`
	err := cr.db.QueryRow(sqlStatement, id).Scan(
		&level.ID,
		&level.Name,
		&level.Permissions,
		&level.CompanyId,
	)

	if err != nil {
		return nil, err
	}

	return &level, nil
}

func (cr *LevelRepositoryMysql) FindAll() ([]*entity.Level, error) {
	sql := `
		SELECT
			id,
			name,
			permissions,
			company_id
		FROM levels
		WHERE deleted_at IS NULL
	`

	rows, err := cr.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var levels []*entity.Level

	for rows.Next() {
		var level entity.Level
		err := rows.Scan(
			&level.ID,
			&level.Name,
			&level.Permissions,
			&level.CompanyId,
		)
		if err != nil {
			return nil, err
		}

		levels = append(levels, &level)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return levels, nil
}

func (cr *LevelRepositoryMysql) Create(level *entity.Level) error {
	sql := `
		INSERT INTO levels (
			id,
			name,
			permissions,
			company_id
		) VALUES (?, ?, ?)
	`

	stmt, err := cr.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		level.ID,
		level.Name,
		level.Permissions,
		level.CompanyId,
	)

	if err != nil {
		return err
	}

	return nil
}

func (cr *LevelRepositoryMysql) Update(level *entity.Level) error {
	sql := `
		UPDATE
			levels
		SET
			name = ?,
			permissions = ?,
			company_id = ?,
			updated_at = NOW()
		WHERE
			id = ?
	`

	stmt, err := cr.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		level.Name,
		level.Permissions,
		level.CompanyId,
		level.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (cr *LevelRepositoryMysql) Delete(id string) error {
	sql := `
		UPDATE
			levels
		SET
			deleted_at = NOW()
		WHERE
			id = ?
	`

	stmt, err := cr.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func (cr *LevelRepositoryMysql) FindUsersByLevelId(levelId string) ([]*entity.User, error) {
	sql := `
		SELECT
			id,
			name,
			email,
			phone,
			level_id,
			company_id 
		FROM users 
		WHERE level_id = ? 
			AND deleted_at IS NULL
	`
	rows, err := cr.db.Query(sql, levelId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Phone,
			&user.LevelId,
			&user.CompanyId,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
