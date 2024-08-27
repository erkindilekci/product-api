package srvc

import (
	"fmt"
	"github.com/erkindilekci/product-app/pkg/domain"
	"github.com/erkindilekci/product-app/pkg/repository"
)

type FakeProductRepository struct {
	products []domain.Product
}

func NewFakeProductRepository(initialProducts []domain.Product) repository.IProductRepository {
	return &FakeProductRepository{initialProducts}
}

func (repository *FakeProductRepository) GetAllProducts() []domain.Product {
	return repository.products
}

func (repository *FakeProductRepository) GetProductsByStore(store string) []domain.Product {
	var products []domain.Product
	for _, product := range repository.products {
		if product.Store == store {
			products = append(products, product)
		}
	}
	return products
}

func (repository *FakeProductRepository) AddProduct(product domain.Product) error {
	product.Id = int64(len(repository.products))
	repository.products = append(repository.products, product)
	return nil
}

func (repository *FakeProductRepository) GetProductById(productId int64) (domain.Product, error) {
	for _, product := range repository.products {
		if product.Id == productId {
			return product, nil
		}
	}
	return domain.Product{}, fmt.Errorf("no product found with the id %d", productId)
}

func (repository *FakeProductRepository) DeleteProductById(productId int64) error {
	for i, product := range repository.products {
		if product.Id == productId {
			repository.products = append(repository.products[:i], repository.products[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("product with id %d not found in repository", productId)
}

func (repository *FakeProductRepository) UpdatePriceById(productId int64, newPrice float32) error {
	for i, product := range repository.products {
		if product.Id == productId {
			repository.products[i].Price = newPrice
			return nil
		}
	}
	return fmt.Errorf("no product found with the id %d", productId)
}
