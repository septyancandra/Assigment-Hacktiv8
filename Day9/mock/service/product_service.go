package service

import (
	"errors"
	"mock/entity"
	"mock/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id string) (*entity.Product, error) {
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (service ProductService) GetAllProducts() ([]*entity.Product, error) {
	products := service.Repository.FindAll()
	if products == nil {
		return nil, errors.New("no products found")
	}
	return products, nil
}
