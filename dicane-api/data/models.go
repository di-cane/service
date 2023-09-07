package data

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/lib/pq"
)

const dbTimeout = time.Second * 3

var db *sql.DB

// New is the function used to create an instance of the data package. It returns the type
// Model, which embeds all the types we want to be available to our application.
func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		Sale: Sale{},
	}
}

type Models struct {
	Sale Sale
}

type Sale struct {
	Sale_id                    string         `json:"sale_id"`
	Father_id                  sql.NullString `json:"father_id"`
	Mother_id                  sql.NullString `json:"mother_id"`
	Is_litter                  bool           `json:"is_litter"`
	Litter_expected_birth_date time.Time      `json:"litter_expected_birth_date"`
	Litter_expected_amount     int            `json:"litter_expected_amount"`
	Litter_confirmed_amount    int            `json:"litter_confirmed_amount"`
	Shipping_age               int            `json:"shipping_age"`
	Birth_date                 time.Time      `json:"birth_date"`
	Breed                      string         `json:"breed"`
	Shipping                   string         `json:"shipping"`
	Vaccines                   []string       `json:"vaccines"`
	Microchip                  bool           `json:"microchip"`
	Pedigree                   string         `json:"pedigree"`
	Weight                     float64        `json:"weight"`
	Height                     float64        `json:"height"`
	Color                      string         `json:"color"`
	Gender                     string         `json:"gender"`
	Traits                     []string       `json:"traits"`
	Adult_max_height           float64        `json:"adult_max_height"`
	Adult_max_weight           float64        `json:"adult_max_weight"`
	Adult_min_height           float64        `json:"adult_min_height"`
	Adult_min_weight           float64        `json:"adult_min_weight"`
	Images                     []string       `json:"images"`
	Price                      float64        `json:"price"`
	Breeder_id                 string         `json:"breeder_id"`
}

// GetAll returns a slice of all users, sorted by last name
func (s *Sale) GetAll() ([]*Sale, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "SELECT * FROM sales"

	rows, err := db.QueryContext(ctx, query)
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
			&sale.Shipping_age,
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

// GetOne returns one sale by id
func (s *Sale) GetOne(id string) (*Sale, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "SELECT * FROM sales WHERE sale_id = $1"

	var sale Sale
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&sale.Sale_id,
		&sale.Father_id,
		&sale.Mother_id,
		&sale.Is_litter,
		&sale.Litter_expected_birth_date,
		&sale.Litter_expected_amount,
		&sale.Litter_confirmed_amount,
		&sale.Shipping_age,
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
		return nil, err
	}

	return &sale, nil
}

// Update updates one sale in the database, using the information
// stored in the receiver sale
func (s *Sale) Update(id string, body Sale) (*Sale, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "UPDATE sales SET" +
		" is_litter = $1," +
		" litter_expected_birth_date = $2," +
		" litter_expected_amount = $3," +
		" litter_confirmed_amount = $4," +
		" shipping_age = $5," +
		" birth_date = $6," +
		" breed = $7," +
		" shipping = $8," +
		" vaccines = $9," +
		" microchip = $10," +
		" pedigree = $11," +
		" weight = $12," +
		" height = $13," +
		" color = $14," +
		" gender = $15," +
		" traits = $16," +
		" adult_max_height = $17," +
		" adult_max_weight = $18," +
		" adult_min_height = $19," +
		" adult_min_weight = $20," +
		" images = $21" +
		" price = $22" +
		" WHERE sale_id = $23"

	_, err := db.ExecContext(ctx, query,
		body.Is_litter,
		body.Litter_expected_birth_date,
		body.Litter_expected_amount,
		body.Litter_confirmed_amount,
		body.Shipping_age,
		body.Birth_date,
		body.Breed,
		body.Shipping,
		pq.Array(body.Vaccines),
		body.Microchip,
		body.Pedigree,
		body.Weight,
		body.Height,
		body.Color,
		body.Gender,
		pq.Array(body.Traits),
		body.Adult_max_height,
		body.Adult_max_weight,
		body.Adult_min_height,
		body.Adult_min_weight,
		pq.Array(body.Images),
		body.Price,
		body.Sale_id,
	)

	if err != nil {
		return nil, err
	}

	return &body, nil
}

// Delete deletes one sale from the database, by Sale.Sale_id
func (s *Sale) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from users where sale_id = $1`

	_, err := db.ExecContext(ctx, stmt, s.Sale_id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByID deletes one user from the database, by ID
func (s *Sale) DeleteByID(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from sales where sale_id = $1`

	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}

// Insert inserts a new sale into the database, and returns the ID of the newly inserted row
func (s *Sale) Insert(sale Sale) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := "INSERT INTO sales (" +
		" is_litter," +
		" litter_expected_birth_date," +
		" litter_expected_amount," +
		" litter_confirmed_amount," +
		" shipping_age," +
		" birth_date," +
		" breed," +
		" shipping," +
		" vaccines," +
		" microchip," +
		" pedigree," +
		" weight," +
		" height," +
		" color," +
		" gender," +
		" traits," +
		" adult_max_height," +
		" adult_max_weight," +
		" adult_min_height," +
		" adult_min_weight," +
		" images," +
		" price )" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)" +
		" RETURNING sale_id"

	var newID string
	err := db.QueryRowContext(ctx, stmt,
		sale.Is_litter,
		sale.Litter_expected_birth_date,
		sale.Litter_expected_amount,
		sale.Litter_confirmed_amount,
		sale.Shipping_age,
		sale.Birth_date,
		sale.Breed,
		sale.Shipping,
		pq.Array(sale.Vaccines),
		sale.Microchip,
		sale.Pedigree,
		sale.Weight,
		sale.Height,
		sale.Color,
		sale.Gender,
		pq.Array(sale.Traits),
		sale.Adult_max_height,
		sale.Adult_max_weight,
		sale.Adult_min_height,
		sale.Adult_min_weight,
		pq.Array(sale.Images),
		sale.Price).Scan(&newID)

	if err != nil {
		return "", err
	}

	return newID, nil
}
