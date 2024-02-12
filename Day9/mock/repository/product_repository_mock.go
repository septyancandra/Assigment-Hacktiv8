package repository

import (
	"github.com/stretchr/testify/mock"
	"mock/entity"
)

type ProductRepositoryMock struct {
	Moc mock.Mock
}

func (repository *ProductRepositoryMock) FindById(id string) *entity.Product {
	arguments := repository.Moc.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}
	product := arguments.Get(0).(entity.Product)
	return &product
}

func (repository *ProductRepositoryMock) FindAll() []*entity.Product {
	arguments := repository.Moc.Called()

	if arguments.Get(0) == nil {
		return nil
	}
	products := arguments.Get(0).([]*entity.Product)
	return products
}
