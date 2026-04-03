package product

import "first-program/domain"

type service struct {
	prdRepo ProductRepo
}

func NewService(prdRepo ProductRepo) Service {
	return &service{
		prdRepo: prdRepo,
	}
}

func (s *service) Create(p domain.Product) (*domain.Product, error) { // call the Create method of the product repository to create a new product in the database
	return s.prdRepo.Create(p)
}

func (s *service) Get(productID int) (*domain.Product, error) { // call the Get method of the product repository to get a product from the database by its ID
	return s.prdRepo.Get(productID)
}

func (s *service) List() ([]*domain.Product, error) { // call the List method of the product repository to get a list of all products from the database
	return s.prdRepo.List()
}

func (s *service) Update(p domain.Product) (*domain.Product, error) { // call the Update method of the product repository to update a product in the database
	return s.prdRepo.Update(p)
}

func (s *service) Delete(productID int) error { // call the Delete method of the product repository to delete a product from the database by its ID
	return s.prdRepo.Delete(productID)
}
