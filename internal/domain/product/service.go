package product

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// create a new product
func (s *Service) CreateProduct(ctx context.Context, req *CreateProductRequest) (*Product, error) {
	product := &Product{
		ID:          uuid.New(),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Active:      true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := s.repo.Create(ctx, product)
	return product, err
}

// get a product by id
func (s *Service) GetProduct(ctx context.Context, id uuid.UUID) (*Product, error) {
	return s.repo.GetByID(ctx, id)
}

// get all products
func (s *Service) GetAllProducts(ctx context.Context) ([]*Product, error) {
	return s.repo.GetAll(ctx)
}

// update a product
func (s *Service) UpdateProduct(ctx context.Context, id uuid.UUID, req *UpdateProductRequest) (*Product, error) {
	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		product.Name = *req.Name
	}
	if req.Description != nil {
		product.Description = *req.Description
	}
	if req.Price != nil {
		product.Price = *req.Price
	}
	if req.Stock != nil {
		product.Stock = *req.Stock
	}
	product.UpdatedAt = time.Now()

	err = s.repo.Update(ctx, id, product)

	return product, err
}

func (s *Service) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	product.Active = false

	updateErr := s.repo.Update(ctx, id, product)
	return updateErr
}

func (s *Service) PermanentDelete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
