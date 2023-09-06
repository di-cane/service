package data

import "context"

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

// Insert inserts a new sale into the database, and returns the ID of the newly inserted row
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
