package repository

import (
	"go-crud/entities"

	"github.com/jmoiron/sqlx"
)

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepositoryDB(db *sqlx.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) FindAll() (*[]entities.Product, error) {
	products := []entities.Product{}
	query := "select * from products"
	if err := r.db.Select(&products, query); err != nil {
		return nil, err
	}

	return &products, nil
}

func (r *productRepository) FindByID(productID int) (*entities.Product, error) {
	product := entities.Product{}
	query := "select * from products where id = ?"
	if err := r.db.Get(&product, query, productID); err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) Insert(product *entities.Product) (*entities.Product, error) {
	query := "insert into products (name, description, price) values (?,?,?)"
	result, err := r.db.Exec(query, product.Name, product.Description, product.Price)
	if err != nil {
		return nil, err
	}

	lastProductID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	product.ID = int(lastProductID)
	return product, nil
}

func (r *productRepository) UpdateByID(productID int, product *entities.Product) (*entities.Product, error) {
	query := "update products set name = ?, description = ?, price = ? where id = ?"
	_, err := r.db.Exec(query, product.Name, product.Description, product.Price, productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *productRepository) DeleteByID(productID int) error {
	query := "delete from  products where id = ?"
	_, err := r.db.Exec(query, productID)
	if err != nil {
		return err
	}

	return nil
}
