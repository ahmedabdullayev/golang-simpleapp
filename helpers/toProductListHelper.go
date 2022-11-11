package helpers

import (
	"database/sql"
	"shop/model"
)

func ToProductListHelper(rows *sql.Rows) ([]model.Product, error) {
	product := model.Product{}
	var products []model.Product

	for rows.Next() {
		var id int
		var title, description string

		err := rows.Scan(&id, &title, &description)

		if err != nil {
			return products, err
		}

		product.Id = id
		product.Title = title
		product.Description = description

		products = append(products, product)
	}

	return products, nil
}
