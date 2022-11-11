package contracts

import "shop/model"

type IProduct interface {
	GetAll() ([]model.Product, error)
	GetOneById(id int) (model.Product, error)
	Create(product *model.Product) error
	UpdateOne(product *model.Product) error
	Delete(id int) error
}
