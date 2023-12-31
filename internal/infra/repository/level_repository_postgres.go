package repository

import (
	"database/sql"
	"strings"

	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type LevelRepositoryPostgres struct {
	db *sql.DB
}

func NewLevelRepositoryPostgres(db *sql.DB) *LevelRepositoryPostgres {
	return &LevelRepositoryPostgres{
		db: db,
	}
}

func (cr *LevelRepositoryPostgres) FindById(id string) (*entity.Level, error) {
	var level entity.Level

	sqlStatement := `
		SELECT
			id,
			name,
			permissions,
			company_id
		FROM levels
		WHERE id = $1
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

func (cr *LevelRepositoryPostgres) FindAll(companyID string) ([]*entity.Level, error) {
	sql := `
		SELECT
			id,
			name,
			permissions,
			company_id
		FROM levels
		WHERE company_id = $1
			AND deleted_at IS NULL
	`

	rows, err := cr.db.Query(sql, companyID)
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

func (cr *LevelRepositoryPostgres) Create(level *entity.Level) error {
	sql := `
		INSERT INTO levels (
			id,
			name,
			permissions,
			company_id,
			created_at,
			updated_at
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			NOW(),
			NOW()
		)
	`

	stmt, err := cr.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		level.ID,
		level.Name,
		strings.Join(level.Permissions, ", "),
		level.CompanyId,
	)

	if err != nil {
		return err
	}

	return nil
}

func (cr *LevelRepositoryPostgres) Update(level *entity.Level) error {
	sql := `
		UPDATE
			levels
		SET
			name = $1,
			permissions = $2,
			company_id = $3,
			updated_at = NOW()
		WHERE
			id = $4
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

func (cr *LevelRepositoryPostgres) Delete(id string) error {
	sql := `
		UPDATE
			levels
		SET
			deleted_at = NOW()
		WHERE
			id = $1
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

func (cr *LevelRepositoryPostgres) FindUsersByLevelId(levelId string) ([]*entity.User, error) {
	sql := `
		SELECT
			id,
			name,
			email,
			phone,
			level_id,
			company_id 
		FROM users 
		WHERE level_id = $1
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
