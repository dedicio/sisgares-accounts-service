package repository

import (
	"database/sql"

	"github.com/dedicio/sisgares-accounts-service/internal/entity"
	"github.com/dedicio/sisgares-accounts-service/pkg/utils"
)

type UserRepositoryPostgres struct {
	db *sql.DB
}

func NewUserRepositoryPostgres(db *sql.DB) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{
		db: db,
	}
}

func (pr *UserRepositoryPostgres) FindById(id string) (*entity.User, error) {
	var user entity.User

	sqlStatement := `
		SELECT
			id,
			name,
			email,
			phone,
			level_id,
			company_id
		FROM users
		WHERE id = ?
			AND deleted_at IS NULL
	`
	err := pr.db.QueryRow(sqlStatement, id).Scan(
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

	return &user, nil
}

func (pr *UserRepositoryPostgres) FindAll() ([]*entity.User, error) {
	sql := `
		SELECT
			id,
			name,
			email,
			phone,
			level_id,
			company_id 
		FROM users 
		WHERE deleted_at IS NULL
	`
	rows, err := pr.db.Query(sql)
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

func (pr *UserRepositoryPostgres) Create(user *entity.User) error {
	password, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	sql := `
		INSERT INTO
			users (
				id,
				name,
				email,
				phone,
				password,
				level_id,
				company_id,
				created_at,
				updated_at
			)
		VALUES (
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			NOW(),
			NOW()
		)
	`
	_, err = pr.db.Exec(
		sql,
		user.ID,
		user.Name,
		user.Email,
		user.Phone,
		password,
		user.LevelId,
		user.CompanyId,
	)

	if err != nil {
		return err
	}

	return nil
}

func (pr *UserRepositoryPostgres) Update(user *entity.User) error {
	sql := `
		UPDATE users
		SET
			name = ?,
			email = ?,
			phone = ?,
			level_id = ?,
			company_id = ?
		WHERE
			id = ?
	`
	_, err := pr.db.Exec(
		sql,
		user.Name,
		user.Email,
		user.Phone,
		user.LevelId,
		user.CompanyId,
		user.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (pr *UserRepositoryPostgres) Delete(id string) error {
	sql := `
		UPDATE users
		SET deleted_at = NOW()
		WHERE id = ?
	`
	_, err := pr.db.Exec(sql, id)

	if err != nil {
		return err
	}

	return nil
}

func (pr *UserRepositoryPostgres) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	sqlStatement := `
		SELECT
			id,
			email,
			password,
			company_id
		FROM users
		WHERE email = ?
			AND deleted_at IS NULL
	`
	err := pr.db.QueryRow(sqlStatement, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CompanyId,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
