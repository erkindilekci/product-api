package srvc

import (
	"github.com/erkindilekci/product-app/pkg/domain"
	"github.com/erkindilekci/product-app/pkg/service"
	"github.com/erkindilekci/product-app/pkg/service/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

var productService service.IProductService

func TestMain(m *testing.M) {
	initialData := []domain.Product{
		{1, "XBOX Series X", 1000.0, 10.0, "Microsoft"},
		{2, "Steelseries Rival 500", 100.0, 20.0, "Amazon"},
		{3, "Asus Vivobook", 600.0, 15.0, "Asus Store"},
		{4, "Macbook Pro M3 Pro", 3000.0, 0.0, "Apple"},
	}
	fakeRepo := NewFakeProductRepository(initialData)
	productService = service.NewProductService(fakeRepo)

	m.Run()
}

func TestGetAllProducts(t *testing.T) {
	actualProducts := productService.GetAllProducts()
	assert.Equal(t, 4, len(actualProducts))
}

func TestAddProduct(t *testing.T) {
	t.Run("ValidProduct", func(t *testing.T) {
		productCreate := dto.ProductCreate{
			Name:     "PlayStation 5",
			Price:    500.0,
			Discount: 5.0,
			Store:    "Sony",
		}
		err := productService.Add(productCreate)
		assert.Nil(t, err)
	})

	t.Run("InvalidProduct", func(t *testing.T) {
		productCreate := dto.ProductCreate{
			Name:     "",
			Price:    -100.0,
			Discount: -10.0,
			Store:    "",
		}
		err := productService.Add(productCreate)
		assert.NotNil(t, err)
	})
}

func TestGetAllProductsByStore(t *testing.T) {
	t.Run("ValidStore", func(t *testing.T) {
		products := productService.GetProductsByStore("Microsoft")
		assert.Equal(t, 1, len(products))
	})

	t.Run("InvalidStore", func(t *testing.T) {
		products := productService.GetProductsByStore("NonExistentStore")
		assert.Equal(t, 0, len(products))
	})
}

func TestGetById(t *testing.T) {
	t.Run("ValidId", func(t *testing.T) {
		product, err := productService.GetById(1)
		assert.Nil(t, err)
		assert.Equal(t, "XBOX Series X", product.Name)
	})

	t.Run("InvalidId", func(t *testing.T) {
		_, err := productService.GetById(999)
		assert.NotNil(t, err)
	})
}

func TestDeleteById(t *testing.T) {
	t.Run("ValidId", func(t *testing.T) {
		err := productService.DeleteById(1)
		assert.Nil(t, err)
	})

	t.Run("InvalidId", func(t *testing.T) {
		err := productService.DeleteById(999)
		assert.NotNil(t, err)
	})
}

func TestUpdatePrice(t *testing.T) {
	t.Run("ValidId", func(t *testing.T) {
		err := productService.UpdatePrice(2, 1200.0)
		assert.Nil(t, err)
	})

	t.Run("InvalidId", func(t *testing.T) {
		err := productService.UpdatePrice(999, 1200.0)
		assert.NotNil(t, err)
	})

	t.Run("InvalidPrice", func(t *testing.T) {
		err := productService.UpdatePrice(2, -100.0)
		assert.NotNil(t, err)
	})
}
