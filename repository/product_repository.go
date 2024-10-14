package repository

import (
	"database/sql"
	"fmt"

	"goapi/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) *ProductRepository {
	return &ProductRepository{
		connection: connection,
	}
}

func (r *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT * FROM product"

	rows, err := r.connection.Query(query)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}

	var products []model.Product
	var product model.Product
	for rows.Next() {
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepository) GetProductByID(productID int) (model.Product, error) {
	var product model.Product

	query, err := r.connection.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		fmt.Println("Error: ", err)
		return model.Product{}, err
	}

	err = query.QueryRow(productID).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (r *ProductRepository) InsertProduct(product model.Product) (int, error) {
	var productID int

	query, err := r.connection.Prepare("INSERT INTO product" +
		"(name, price)" +
		"VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println("Error: ", err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&productID)
	if err != nil {
		fmt.Println("Error: ", err)
		return 0, err
	}

	query.Close()

	return productID, nil
}
