package repository

import (
	"database/sql"

	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type AddressRepositoryPostgres struct {
	db *sql.DB
}

func NewAddressRepositoryPostgres(db *sql.DB) *AddressRepositoryPostgres {
	return &AddressRepositoryPostgres{
		db: db,
	}
}

func (ar *AddressRepositoryPostgres) FindById(id string) (*entity.Address, error) {
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
		WHERE id = $1
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

func (ar *AddressRepositoryPostgres) Create(address *entity.Address) error {
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
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
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

func (ar *AddressRepositoryPostgres) Update(address *entity.Address) error {
	sql := `
		UPDATE
			addresses
		SET
			street = $1,
			number = $2,
			complement = $3,
			neighborhood = $4,
			city = $5,
			state = $6,
			country = $7,
			zip_code = $8,
			company_id = $9,
			updated_at = NOW()
		WHERE
			id = $10
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

func (ar *AddressRepositoryPostgres) Delete(id string) error {
	sql := `
		UPDATE
			addresses
		SET deleted_at = NOW()
		WHERE id = $1
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

func (ar *AddressRepositoryPostgres) DeleteByCompanyId(companyId string) error {
	sql := `
		UPDATE
			addresses
		SET deleted_at = NOW()
		WHERE company_id = $1
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
