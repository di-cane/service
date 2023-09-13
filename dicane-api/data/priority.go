package data

import (
	"context"
	"database/sql"
	"log"
)

type Priority struct {
	Priority_id  string         `json:"priority_id"`
	Sale_id      string         `json:"sale_id"`
	Customer_id  sql.NullString `json:"customer_id"`
	Position     int            `json:"position"`
	Price        float64        `json:"price"`
	Is_available bool           `json:"is_available"`
}

func (p *Priority) GetPriorityList(saleId string) ([]*Priority, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "SELECT * FROM priority WHERE sale_id = $1"

	rows, err := db.QueryContext(ctx, query, saleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var priorities []*Priority

	for rows.Next() {
		var priority Priority
		err := rows.Scan(
			&priority.Priority_id,
			&priority.Sale_id,
			&priority.Customer_id,
			&priority.Position,
			&priority.Price,
			&priority.Is_available,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}
		priorities = append(priorities, &priority)
	}
	return priorities, nil
}

func (p *Priority) InsertPriorityList(priorities []Priority) ([]Priority, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "INSERT INTO priority (sale_id, customer_id, position, price, is_available) VALUES ($1, $2, $3, $4, $5) RETURNING *"

	for _, priority := range priorities {
		err := db.QueryRowContext(ctx, query, priority.Sale_id, priority.Customer_id, priority.Position, priority.Price, priority.Is_available).Scan(
			&priority.Priority_id,
			&priority.Sale_id,
			&priority.Customer_id,
			&priority.Position,
			&priority.Price,
			&priority.Is_available,
		)
		if err != nil {
			return nil, err
		}
	}

	return priorities, nil
}
