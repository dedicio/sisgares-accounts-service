package repository

import (
	"database/sql"

	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type AddressRepositoryMysql struct {
	db *sql.DB
}

func NewAddressRepositoryMysql(db *sql.DB) *AddressRepositoryMysql {
	return &AddressRepositoryMysql{
		db: db,
	}
}

func (ar *AddressRepositoryMysql) FindById(id string) (*entity.Address, error) {
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
		WHERE id = ?
			AND deleted_at IS NULL
	`
	err := ar.db.QueryRow(sqlStatement, id).Scan(
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

func (ar *AddressRepositoryMysql) Create(address *entity.Address) error {
	sql := `
		INSERT INTO addresses (
			id,
			street,
			number,
			complement,
			neighborhood,
			city,
			state,
			country,
			zip_code,
			company_id,
			created_at,
			updated_at
		) VALUES (
			?,
			?,
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

	stmt, err := ar.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		address.ID,
		address.Street,
		address.Number,
		address.Complement,
		address.Neighborhood,
		address.City,
		address.State,
		address.Country,
		address.ZipCode,
		address.CompanyId,
	)

	if err != nil {
		return err
	}

	return nil
}

func (ar *AddressRepositoryMysql) Update(address *entity.Address) error {
	sql := `
		UPDATE
			addresses
		SET
			street = ?,
			number = ?,
			complement = ?,
			neighborhood = ?,
			city = ?,
			state = ?,
			country = ?,
			zip_code = ?,
			company_id = ?,
			updated_at = NOW()
		WHERE
			id = ?
	`

	stmt, err := ar.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		address.Street,
		address.Number,
		address.Complement,
		address.Neighborhood,
		address.City,
		address.State,
		address.Country,
		address.ZipCode,
		address.CompanyId,
		address.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (ar *AddressRepositoryMysql) Delete(id string) error {
	sql := `
		UPDATE
			addresses
		SET deleted_at = NOW()
		WHERE id = ?
	`

	stmt, err := ar.db.Prepare(sql)
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

func (ar *AddressRepositoryMysql) DeleteByCompanyId(companyId string) error {
	sql := `
		UPDATE
			addresses
		SET deleted_at = NOW()
		WHERE company_id = ?
	`

	stmt, err := ar.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(companyId)

	if err != nil {
		return err
	}

	return nil
}
