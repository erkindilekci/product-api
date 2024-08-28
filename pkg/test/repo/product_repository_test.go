package repo

import (
	"context"
	"fmt"
	"github.com/erkindilekci/product-api/pkg/common/postgresql"
	"github.com/erkindilekci/product-api/pkg/domain"
	"github.com/erkindilekci/product-api/pkg/repository"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var productRepo repository.IProductRepository
var databasePool *pgxpool.Pool
var testContext context.Context

func TestMain(m *testing.M) {
	testContext = context.Background()

	databasePool = postgresql.GetConnectionPool(testContext, postgresql.Config{
		Host:                  "localhost",
		Port:                  "5433",
		UserName:              "postgres",
		Password:              "password",
		DbName:                "productapp",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	})

	productRepo = repository.NewProductRepository(databasePool)

	fmt.Println("Before / Setup")
	exitCode := m.Run()

	fmt.Println("After / Teardown")
	os.Exit(exitCode)
}

func setupTestData(ctx context.Context, dbPool *pgxpool.Pool) {
	TestDataInitialize(ctx, dbPool)
}

func teardownTestData(ctx context.Context, dbPool *pgxpool.Pool) {
	TruncateTestData(ctx, dbPool)
}

func TestGetAllProducts(t *testing.T) {
	setupTestData(testContext, databasePool)

	actualProducts := productRepo.GetAllProducts()

	t.Run("TestGetAllProductsLength", func(t *testing.T) {
		actualLength := len(actualProducts)
		expectedLength := 4

		assert.Equal(t, expectedLength, actualLength)
	})

	t.Run("TestGetAllProductsContent", func(t *testing.T) {
		expectedProducts := []domain.Product{
			{1, "XBOX Series X", 1000.0, 10.0, "Microsoft"},
			{2, "Steelseries Rival 500", 100.0, 20.0, "Amazon"},
			{3, "Asus Vivobook", 600.0, 15.0, "Asus Store"},
			{4, "Macbook Pro M3 Pro", 3000.0, 0.0, "Apple"},
		}
		assert.Equal(t, expectedProducts, actualProducts)
	})

	teardownTestData(testContext, databasePool)
}

func TestGetAllProductsByStore(t *testing.T) {
	setupTestData(testContext, databasePool)

	actualProducts := productRepo.GetProductsByStore("Apple")
	expectedProducts := []domain.Product{
		{4, "Macbook Pro M3 Pro", 3000.0, 0.0, "Apple"},
	}

	t.Run("TestGetAllProductsByStoreLength", func(t *testing.T) {
		actualLength := len(actualProducts)
		expectedLength := 1

		assert.Equal(t, expectedLength, actualLength)
	})

	t.Run("TestGetAllProductsByStoreContent", func(t *testing.T) {
		assert.Equal(t, expectedProducts, actualProducts)
	})

	teardownTestData(testContext, databasePool)
}

func TestAddProduct(t *testing.T) {
	newProduct := domain.Product{Name: "Product 1", Price: 100.0, Discount: 20.0, Store: "Store 1"}
	err := productRepo.AddProduct(newProduct)
	allProducts := productRepo.GetAllProducts()

	t.Run("TestAddProductNoError", func(t *testing.T) {
		assert.NoError(t, err)
	})

	t.Run("TestAddProductLength", func(t *testing.T) {
		assert.Equal(t, 1, len(allProducts))
	})

	t.Run("TestAddProductContent", func(t *testing.T) {
		addedProduct := allProducts[0]
		expectedProduct := domain.Product{Id: 1, Name: "Product 1", Price: 100.0, Discount: 20.0, Store: "Store 1"}
		assert.Equal(t, expectedProduct, addedProduct)
	})

	teardownTestData(testContext, databasePool)
}

func TestGetProductById(t *testing.T) {
	setupTestData(testContext, databasePool)

	t.Run("TestGetProductByIdValid", func(t *testing.T) {
		product, err := productRepo.GetProductById(1)
		expectedProduct := domain.Product{Id: 1, Name: "XBOX Series X", Price: 1000.0, Discount: 10.0, Store: "Microsoft"}

		assert.NoError(t, err)
		assert.Equal(t, expectedProduct, product)
	})

	teardownTestData(testContext, databasePool)
}

func TestDeleteProductById(t *testing.T) {
	setupTestData(testContext, databasePool)

	t.Run("TestDeleteProductByIdValid", func(t *testing.T) {
		err := productRepo.DeleteProductById(4)
		assert.NoError(t, err)
	})

	t.Run("TestDeleteProductByIdContent", func(t *testing.T) {
		deletedProduct := domain.Product{Id: 4, Name: "Macbook Pro M3 Pro", Price: 3000.0, Store: "Apple"}
		assert.NotContains(t, productRepo.GetAllProducts(), deletedProduct)
	})

	teardownTestData(testContext, databasePool)
}

func TestUpdatePriceById(t *testing.T) {
	setupTestData(testContext, databasePool)

	t.Run("TestUpdatePriceByIdValid", func(t *testing.T) {
		err := productRepo.UpdatePriceById(4, 3200.0)
		assert.NoError(t, err)
	})

	t.Run("TestUpdatePriceByIdContent", func(t *testing.T) {
		updatedProduct := domain.Product{Id: 4, Name: "Macbook Pro M3 Pro", Price: 3200.0, Store: "Apple"}
		assert.Contains(t, productRepo.GetAllProducts(), updatedProduct)
	})

	teardownTestData(testContext, databasePool)
}
