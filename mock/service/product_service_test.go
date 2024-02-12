package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mock/entity"
	"mock/repository"
	"testing"
)

var productRepostitory = &repository.ProductRepositoryMock{Moc: mock.Mock{}}
var productService = ProductService{Repository: productRepostitory}

func TestProductService_GetOneProductNotFound(t *testing.T) {
	productRepostitory.Moc.On("FindById", "1").Return(nil)
	product, err := productService.GetOneProduct("1")
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "Error Response has to be product not found")
}

func TestProductService_GetOneProduct(t *testing.T) {
	product := entity.Product{
		Id:   "2",
		Name: "kaca mata",
	}
	productRepostitory.Moc.On("FindById", "2").Return(product)

	result, err := productService.GetOneProduct("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, product.Id, result.Id, "result has to be '2' ")
	assert.Equal(t, product.Name, result.Name, "result has to be 'kaca mata' ")
	assert.Equal(t, &product, result, "result has to be a product data with id '2' ")

}

func TestFindAllProducts(t *testing.T) {
	products := []*entity.Product{
		{Id: "1", Name: "Product 1"},
		{Id: "2", Name: "Product 2"},
	}
	productRepostitory.Moc.On("FindAll").Return(products)

	result, err := productService.GetAllProducts()

	assert.Nil(t, err)

	assert.Equal(t, products, result)
}

func TestFindAllProductsFail(t *testing.T) {
	products := []*entity.Product{
		{Id: "1", Name: "Product 1"},
		{Id: "2", Name: "Product 2"},
	}
	productRepostitory.Moc.On("FindAll").Return(products)

	result, err := productService.GetAllProducts()

	assert.Nil(t, err)

	assert.NotEqual(t, len(products), len(result)+1)
}
