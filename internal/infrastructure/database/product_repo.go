package database

import (
	"context"
	"database/sql"
	"go_commerce/internal/domain/product"

	"github.com/google/uuid"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(ctx context.Context, p *product.Product) error {
	query := `
		INSERT INTO products (id, name, description, price, stock, active, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := r.db.ExecContext(ctx, query, p.ID, p.Name, p.Description, p.Price, p.Stock, p.Active, p.CreatedAt, p.UpdatedAt)

	return err
}

func (r *ProductRepository) GetByID(ctx context.Context, id uuid.UUID) (*product.Product, error) {
	query := `
		SELECT id, name, description, price, stock, created_at, updated_at FROM products WHERE id = $1 AND active = $2
	`
	p := &product.Product{}
	err := r.db.QueryRowContext(ctx, query, id, true).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CreatedAt, &p.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return p, err
}

func (r *ProductRepository) GetAll(ctx context.Context) ([]*product.Product, error) {
	query := `
		SELECT id, name, description, price, stock, created_at, updated_at FROM products WHERE active = $1
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*product.Product
	for rows.Next() {
		p := &product.Product{}
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CreatedAt, &p.UpdatedAt)

		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *ProductRepository) Update(ctx context.Context, id uuid.UUID, p *product.Product) error {
	query := `
		UPDATE products
		SET name=$2, description=$3, price=$4, stock=$5, updated_at=$6, active=$7
		WHERE id=$1 AND active=$8
	`

	_, err := r.db.ExecContext(ctx, query, id, p.Name, p.Description, p.Price, p.Stock, p.UpdatedAt, p.Active, true)
	return err
}

func (r *ProductRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `
		DELETE FROM products WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
