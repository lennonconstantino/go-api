package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository interface {
	GetProducts() ([]model.Product, error)
	CreateProduct(product model.Product) (int, error)
	GetProductById(id_product int) (*model.Product, error)
	DeleteProduct(id_product int) error
	UpdateProduct(id_product int, product model.Product) error
}

type ProductRepositoryImpl struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		connection: connection,
	}
}

func (pr ProductRepositoryImpl) GetProducts() ([]model.Product, error) {
	query := "SELECT id, name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr ProductRepositoryImpl) CreateProduct(product model.Product) (int, error) {
	var id int
	query, err := pr.connection.Prepare("INSERT INTO product" +
		"(name, price)" +
		" VALUES ($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}

func (pr ProductRepositoryImpl) GetProductById(id_product int) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var product model.Product

	err = query.QueryRow(id_product).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &product, nil
}

func (pr ProductRepositoryImpl) DeleteProduct(id_product int) error {
	statement, err := pr.connection.Prepare("DELETE FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(id_product); err != nil {
		return err
	}

	return nil
}

func (pr ProductRepositoryImpl) UpdateProduct(id_product int, product model.Product) error {
	statement, err := pr.connection.Prepare(
		"update product set price = $1, name = $2 where id = $3",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(product.Price, product.Name, id_product); err != nil {
		return err
	}

	return nil
}
