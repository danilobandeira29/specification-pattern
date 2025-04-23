package v2

import (
	"database/sql"
	"fmt"
	"github.com/danilobandeira29/specification-pattern/repository/v2/specification"
	"log"
)

type Product struct {
	ID       string
	Name     string
	Category string
	Price    float64
	IsActive bool
}

type Repository interface {
	FindBy(s specification.Specification) ([]Product, error)
}

type ProductRepository struct {
	conn *sql.DB
}

func (p *ProductRepository) FindBy(s specification.Specification) ([]Product, error) {
	query, args := s.ToSQL()
	q := "select id, name, category, price, is_active from products where " + query
	rows, err := p.conn.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("find by: %v\nquery: %s", err, q)
	}
	log.Println(q, args)
	var result []Product
	for rows.Next() {
		var product Product
		errScan := rows.Scan(&product.ID, &product.Name, &product.Category, &product.Price, &product.IsActive)
		if errScan != nil {
			return nil, fmt.Errorf("error scan: %v", errScan)
		}
		result = append(result, product)
	}
	return result, nil
}

func New(conn *sql.DB) Repository {
	return &ProductRepository{
		conn: conn,
	}
}
