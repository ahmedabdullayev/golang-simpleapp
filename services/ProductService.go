package services

import (
	"shop/contracts"
	"shop/model"
)

type ProductService struct {
	repository contracts.IProduct
}

func NewProductService(repository contracts.IProduct) contracts.IProduct {
	return &ProductService{
		repository: repository,
	}
}

func (service ProductService) GetAll() ([]model.Product, error) {
	return service.repository.GetAll()
}

func (service ProductService) GetOneById(id int) (model.Product, error) {
	return service.repository.GetOneById(id)
}

func (service ProductService) Create(product *model.Product) error {
	return service.repository.Create(product)
}

func (service ProductService) UpdateOne(product *model.Product) error {
	return service.repository.UpdateOne(product)
}

func (service ProductService) Delete(id int) error {
	return service.repository.Delete(id)
}
