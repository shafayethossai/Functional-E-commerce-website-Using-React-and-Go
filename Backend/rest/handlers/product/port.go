package product

import "first-program/domain"

type Service interface {
	Create(domain.Product) (*domain.Product, error)
	Get(productID int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Update(p domain.Product) (*domain.Product, error)
	Delete(productID int) error
}
