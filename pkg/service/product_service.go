package service

import (
	"errors"
	"github.com/erkindilekci/product-app/pkg/domain"
	"github.com/erkindilekci/product-app/pkg/repository"
	"github.com/erkindilekci/product-app/pkg/service/dto"
)

type IProductService interface {
	Add(productCreate dto.ProductCreate) error
	GetAllProducts() []domain.Product
	GetProductsByStore(store string) []domain.Product
	GetById(productId int64) (domain.Product, error)
	DeleteById(productId int64) error
	UpdatePrice(productId int64, newPrice float32) error
}

type ProductService struct {
	productRepository repository.IProductRepository
}

func NewProductService(productRepository repository.IProductRepository) IProductService {
	return &ProductService{productRepository}
}

func (service *ProductService) Add(productCreate dto.ProductCreate) error {
	err := validateProductCreate(productCreate)
	if err != nil {
		return err
	}

	product := productCreateToProduct(productCreate)
	return service.productRepository.AddProduct(product)
}

func (service *ProductService) GetAllProducts() []domain.Product {
	return service.productRepository.GetAllProducts()
}

func (service *ProductService) GetProductsByStore(store string) []domain.Product {
	return service.productRepository.GetProductsByStore(store)
}

func (service *ProductService) GetById(productId int64) (domain.Product, error) {
	return service.productRepository.GetProductById(productId)
}

func (service *ProductService) DeleteById(productId int64) error {
	_, err := service.GetById(productId)
	if err != nil {
		return err
	}

	return service.productRepository.DeleteProductById(productId)
}

func (service *ProductService) UpdatePrice(productId int64, newPrice float32) error {
	_, err := service.GetById(productId)
	if err != nil {
		return err
	}

	if newPrice < 0 {
		return errors.New("price can't be less than zero")
	}

	return service.productRepository.UpdatePriceById(productId, newPrice)
}

func validateProductCreate(productCreate dto.ProductCreate) error {
	if productCreate.Name == "" {
		return errors.New("name can't be empty")
	}
	if productCreate.Price < 0 {
		return errors.New("price can't be less than zero")
	}
	if productCreate.Discount < 0 {
		return errors.New("discount can't be less than zero")
	}
	if productCreate.Store == "" {
		return errors.New("store can't be empty")
	}
	return nil
}

func productCreateToProduct(productCreate dto.ProductCreate) domain.Product {
	return domain.Product{
		Name:     productCreate.Name,
		Price:    productCreate.Price,
		Discount: productCreate.Discount,
		Store:    productCreate.Store,
	}
}
