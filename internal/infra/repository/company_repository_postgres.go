package repository

import (
	"database/sql"
	"fmt"

	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type CompanyRepositoryPostgres struct {
	db *sql.DB
}

func NewCompanyRepositoryPostgres(db *sql.DB) *CompanyRepositoryPostgres {
	return &CompanyRepositoryPostgres{
		db: db,
	}
}

func (cr *CompanyRepositoryPostgres) FindById(id string) (*entity.Company, error) {
	var company entity.Company
	var address entity.Address
	fmt.Println("chega no repository", company)

	sqlStatement := `
		SELECT
			c.id,
			c.name,
			COALESCE(c.document, ''),
			c.status,
			COALESCE(a.street, ''),
			COALESCE(a.number, ''),
			COALESCE(a.complement, ''),
			COALESCE(a.neighborhood, ''),
			COALESCE(a.city, ''),
			COALESCE(a.state, ''),
			COALESCE(a.country, ''),
			COALESCE(a.zip_code, '')
		FROM companies c
		LEFT JOIN addresses a
			ON a.company_id = c.id
		WHERE c.id = $1
			AND c.deleted_at IS NULL
	`
	err := cr.db.QueryRow(sqlStatement, id).Scan(
		&company.ID,
		&company.Name,
		&company.Document,
		&company.Status,
		&address.Street,
		&address.Number,
		&address.Complement,
		&address.Neighborhood,
		&address.City,
		&address.State,
		&address.Country,
		&address.ZipCode,
	)

	company.Address = &address

	if err != nil {
		return nil, err
	}

	fmt.Println("repository", company)

	return &company, nil
}

func (cr *CompanyRepositoryPostgres) FindAll() ([]*entity.Company, error) {
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

	rows, err := cr.db.Query(sql)

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

func (cr *CompanyRepositoryPostgres) Create(company *entity.Company) error {
	sql := `
		INSERT INTO companies (
			id,
			name,
			document,
			status,
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

	_, err := cr.db.Exec(
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

func (cr *CompanyRepositoryPostgres) Update(company *entity.Company) error {
	sql := `
		UPDATE companies, addresses
		SET
			companies.name = $1,
			companies.document = $2,
			companies.status = $3,
			addresses.street = $4,
			addresses.number = $5,
			addresses.complement = $6,
			addresses.neighborhood = $7,
			addresses.city = $8,
			addresses.state = $9,
			addresses.country = $10,
			addresses.zip_code = $11,
			companies.updated_at = NOW(),
			addresses.updated_at = NOW()
		WHERE companies.id = $12
			AND addresses.company_id = $13
	`

	_, err := cr.db.Exec(
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

func (cr *CompanyRepositoryPostgres) Delete(id string) error {
	sql := `
		UPDATE companies, addresses
		SET
			companies.deleted_at = NOW(),
			addresses.deleted_at = NOW()
		WHERE id = $1
			AND addresses.company_id = $2
	`

	_, err := cr.db.Exec(sql, id)

	if err != nil {
		return err
	}

	return nil
}

func (cr *CompanyRepositoryPostgres) FindAddressByCompanyId(companyId string) (*entity.Address, error) {
	var address entity.Address

	sqlStatement := `
		SELECT
			id,
			street,
			number,
			complement,
			neighborhood,
			city,
			state,
			country,
			zip_code,
			company_id
		FROM addresses
		WHERE company_id = $1
			AND deleted_at IS NULL
	`
	err := cr.db.QueryRow(sqlStatement, companyId).Scan(
		&address.ID,
		&address.Street,
		&address.Number,
		&address.Complement,
		&address.Neighborhood,
		&address.City,
		&address.State,
		&address.Country,
		&address.ZipCode,
		&address.CompanyId,
	)

	if err != nil {
		return nil, err
	}

	return &address, nil
}
