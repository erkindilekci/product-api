package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/erkindilekci/product-app/pkg/domain"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type IProductRepository interface {
	GetAllProducts() []domain.Product
	GetProductsByStore(store string) []domain.Product
	AddProduct(product domain.Product) error
	GetProductById(productId int64) (domain.Product, error)
	DeleteProductById(productId int64) error
	UpdatePriceById(productId int64, newPrice float32) error
}

type ProductRepository struct {
	dbPool *pgxpool.Pool
}

func NewProductRepository(dbPool *pgxpool.Pool) IProductRepository {
	return &ProductRepository{dbPool}
}

func (repository *ProductRepository) GetAllProducts() []domain.Product {
	ctx := context.Background()
	productRows, err := repository.dbPool.Query(ctx, "SELECT * FROM products")
	if err != nil {
		log.Errorf("error while getting all products: %v", err)
		return []domain.Product{}
	}

	return extractProductsFromRows(productRows)
}

func (repository *ProductRepository) GetProductsByStore(store string) []domain.Product {
	ctx := context.Background()
	productRows, err := repository.dbPool.Query(ctx, "SELECT * FROM products WHERE store = $1", store)
	if err != nil {
		log.Errorf("error while getting all products by store: %v", err)
		return []domain.Product{}
	}

	return extractProductsFromRows(productRows)
}

func (repository *ProductRepository) AddProduct(product domain.Product) error {
	ctx := context.Background()

	insertStatement := "INSERT INTO products (name, price, discount, store) VALUES ($1, $2, $3, $4)"

	addNewProduct, err := repository.dbPool.Exec(ctx, insertStatement, product.Name, product.Price, product.Discount, product.Store)
	if err != nil {
		log.Errorf("error while adding a new product: %v", err)
		return err
	}

	log.Info(fmt.Sprintf("Product added successfully: %v", addNewProduct))
	return nil
}

func (repository *ProductRepository) GetProductById(productId int64) (domain.Product, error) {
	ctx := context.Background()

	var product domain.Product
	productRow := repository.dbPool.QueryRow(ctx, "SELECT * FROM products WHERE id = $1", productId)

	err := productRow.Scan(&product.Id, &product.Name, &product.Price, &product.Discount, &product.Store)
	if err != nil && err.Error() == "no rows in result set" {
		errStr := fmt.Sprintf("error while getting product by id: %d", productId)
		log.Errorf(errStr)
		return domain.Product{}, errors.New(errStr)
	}

	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (repository *ProductRepository) DeleteProductById(productId int64) error {
	ctx := context.Background()

	_, err := repository.dbPool.Exec(ctx, "DELETE FROM products WHERE id = $1", productId)
	if err != nil {
		return err
	}

	log.Info("Product deleted successfully")

	return nil
}

func (repository *ProductRepository) UpdatePriceById(productId int64, newPrice float32) error {
	ctx := context.Background()

	_, err := repository.dbPool.Exec(ctx, "UPDATE products SET price = $1 WHERE id = $2", newPrice, productId)
	if err != nil {
		return err
	}

	log.Info("Price updated successfully")

	return nil
}

func extractProductsFromRows(productRows pgx.Rows) []domain.Product {
	var products []domain.Product

	for productRows.Next() {
		var product domain.Product
		productRows.Scan(&product.Id, &product.Name, &product.Price, &product.Discount, &product.Store)
		products = append(products, product)
	}

	return products
}
