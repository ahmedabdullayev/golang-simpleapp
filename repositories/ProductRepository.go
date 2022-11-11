package repositories

import (
	"database/sql"
	"shop/contracts"
	"shop/helpers"
	"shop/model"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) contracts.IProduct {
	return &ProductRepository{DB: db}
}

func (rep ProductRepository) GetAll() ([]model.Product, error) {
	var products []model.Product
	selectAllQuery := "SELECT * FROM products ORDER BY id DESC"
	productsQuery, err := rep.DB.Query(selectAllQuery)
	if err != nil {
		return products, err
	}
	products, err = helpers.ToProductListHelper(productsQuery)

	return products, err
}

func (rep ProductRepository) GetOneById(id int) (model.Product, error) {
	product := model.Product{}
	row := rep.DB.QueryRow("SELECT * FROM products WHERE id = $1", id)
	err := row.Scan(&product.Id, &product.Title, &product.Description)
	if err != nil {
		panic(err.Error())
	}

	return product, nil
}

func (rep ProductRepository) Create(product *model.Product) error {
	stmt, err := rep.DB.Prepare("INSERT INTO products (title, description) VALUES($1, $2)")

	if err != nil {
		return err
	}

	stmt.Exec(product.Title, product.Description)

	return nil
}

func (rep ProductRepository) UpdateOne(product *model.Product) error {
	addQuery, err := rep.DB.Prepare("UPDATE products SET title=$1, description=$2 WHERE id=$3")
	if err != nil {
		panic(err.Error())
	}
	addQuery.Exec(product.Title, product.Description, product.Id)

	return nil
}

func (rep ProductRepository) Delete(id int) error {
	deleteQuery, err := rep.DB.Prepare("DELETE FROM products WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}
	deleteQuery.Exec(id)

	return nil
}
