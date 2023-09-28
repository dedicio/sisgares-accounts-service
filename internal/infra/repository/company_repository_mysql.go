package repository

import (
	"database/sql"

	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type CompanyRepositoryMysql struct {
	db *sql.DB
}

func NewCompanyRepositoryMysql(db *sql.DB) *CompanyRepositoryMysql {
	return &CompanyRepositoryMysql{
		db: db,
	}
}

func (pr *CompanyRepositoryMysql) FindById(id string) (*entity.Company, error) {
	var company entity.Company

	sqlStatement := `
		SELECT
			c.id,
			c.name,
			c.document,
			c.status,
			a.street,
			a.number,
			a.complement,
			a.neighborhood,
			a.city,
			a.state,
			a.country,
			a.zip_code
		FROM companies c
		LEFT JOIN addresses a
			ON a.company_id = c.id
		WHERE c.id = ?
			AND c.deleted_at IS NULL
	`
	err := pr.db.QueryRow(sqlStatement, id).Scan(
		&company.ID,
		&company.Name,
		&company.Document,
		&company.Status,
		&company.Address.Street,
		&company.Address.Number,
		&company.Address.Complement,
		&company.Address.Neighborhood,
		&company.Address.City,
		&company.Address.State,
		&company.Address.Country,
		&company.Address.ZipCode,
	)

	if err != nil {
		return nil, err
	}

	return &company, nil
}

func (pr *CompanyRepositoryMysql) FindAll() ([]*entity.Company, error) {
	sql := `
		SELECT
			c.id,
			c.name,
			c.document,
			c.status,
			a.street,
			a.number,
			a.complement,
			a.neighborhood,
			a.city,
			a.state,
			a.country,
			a.zip_code
		FROM companies c
		LEFT JOIN addresses a
			ON a.company_id = c.id
		WHERE c.deleted_at IS NULL
	`

	rows, err := pr.db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	companies := []*entity.Company{}
	for rows.Next() {
		var company entity.Company
		err := rows.Scan(
			&company.ID,
			&company.Name,
			&company.Document,
			&company.Status,
			&company.Address.Street,
			&company.Address.Number,
			&company.Address.Complement,
			&company.Address.Neighborhood,
			&company.Address.City,
			&company.Address.State,
			&company.Address.Country,
			&company.Address.ZipCode,
		)
		if err != nil {
			return nil, err
		}
		companies = append(companies, &company)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}

func (pr *CompanyRepositoryMysql) Create(company *entity.Company) error {
	sql := `
		INSERT INTO companies (
			id,
			name,
			document,
			status,
			created_at,
			updated_at
		) VALUES (
			?,
			?,
			?,
			?,
			NOW(),
			NOW()
		)
	`

	_, err := pr.db.Exec(
		sql,
		company.ID,
		company.Name,
		company.Document,
		company.Status,
	)

	if err != nil {
		return err
	}

	return nil
}

func (pr *CompanyRepositoryMysql) Update(company *entity.Company) error {
	sql := `
		UPDATE companies, addresses
		SET
			companies.name = ?,
			companies.document = ?,
			companies.status = ?,
			addresses.street = ?,
			addresses.number = ?,
			addresses.complement = ?,
			addresses.neighborhood = ?,
			addresses.city = ?,
			addresses.state = ?,
			addresses.country = ?,
			addresses.zip_code = ?,
			companies.updated_at = NOW(),
			addresses.updated_at = NOW()
		WHERE companies.id = ?
			AND addresses.company_id = ?
	`

	_, err := pr.db.Exec(
		sql,
		company.Name,
		company.Document,
		company.Status,
		company.Address.Street,
		company.Address.Number,
		company.Address.Complement,
		company.Address.Neighborhood,
		company.Address.City,
		company.Address.State,
		company.Address.Country,
		company.Address.ZipCode,
		company.ID,
		company.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (pr *CompanyRepositoryMysql) Delete(id string) error {
	sql := `
		UPDATE companies, addresses
		SET
			companies.deleted_at = NOW(),
			addresses.deleted_at = NOW()
		WHERE id = ?
			AND addresses.company_id = ?
	`

	_, err := pr.db.Exec(sql, id)

	if err != nil {
		return err
	}

	return nil
}
