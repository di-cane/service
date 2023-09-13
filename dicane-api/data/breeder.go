package data

import (
	"context"
	"log"

	"github.com/lib/pq"
)

type Breeder struct {
	Breeder_id  string `json:"breeder_id"`
	Kennel_name string `json:"kennel_name"`
	Email       string `json:"email"`
	Cnpj        string `json:"cnpj"`
	Document    string `json:"document"`
	Logo        string `json:"logo"`
}

// GetOne returns one sale by id
func (b *Breeder) GetByEmail(email string) (*Breeder, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "SELECT * FROM breeder WHERE email = $1"

	var breeder Breeder
	row := db.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&breeder.Breeder_id,
		&breeder.Kennel_name,
		&breeder.Email,
		&breeder.Cnpj,
		&breeder.Document,
		&breeder.Logo,
	)

	if err != nil {
		return nil, err
	}

	return &breeder, nil
}

// Insert inserts a new breeder into the database, and returns the ID of the newly inserted row
func (b *Breeder) Insert(breeder Breeder) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := "INSERT INTO breeder (" +
		" kennel_name," +
		" email," +
		" cnpj," +
		" document," +
		" logo )" +
		" VALUES ($1, $2, $3, $4, $5)" +
		" RETURNING breeder_id"

	var newID string
	err := db.QueryRowContext(ctx, stmt,
		breeder.Kennel_name,
		breeder.Email,
		breeder.Cnpj,
		breeder.Document,
		breeder.Logo).Scan(&newID)

	if err != nil {
		return "", err
	}

	return newID, nil
}

// GetAll returns a slice of all users, sorted by last name
func (b *Breeder) GetBreederSales(id string) ([]*Sale, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "SELECT * FROM sales WHERE breeder_id = $1"

	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sales []*Sale

	for rows.Next() {
		var sale Sale
		err := rows.Scan(
			&sale.Sale_id,
			&sale.Father_id,
			&sale.Mother_id,
			&sale.Is_litter,
			&sale.Litter_expected_birth_date,
			&sale.Litter_expected_amount,
			&sale.Litter_confirmed_amount,
			&sale.Shipping_date,
			&sale.Birth_date,
			&sale.Breed,
			&sale.Shipping,
			pq.Array(&sale.Vaccines),
			&sale.Microchip,
			&sale.Pedigree,
			&sale.Weight,
			&sale.Height,
			&sale.Color,
			&sale.Gender,
			pq.Array(&sale.Traits),
			&sale.Adult_max_height,
			&sale.Adult_max_weight,
			&sale.Adult_min_height,
			&sale.Adult_min_weight,
			pq.Array(&sale.Images),
			&sale.Price,
			&sale.Breeder_id,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		sales = append(sales, &sale)
	}

	return sales, nil
}
