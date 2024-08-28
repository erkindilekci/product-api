package main

import (
	"context"
	"github.com/erkindilekci/product-api/pkg/common/app"
	"github.com/erkindilekci/product-api/pkg/common/postgresql"
	"github.com/erkindilekci/product-api/pkg/controller"
	"github.com/erkindilekci/product-api/pkg/repository"
	"github.com/erkindilekci/product-api/pkg/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	ctx := context.Background()
	configurationManager := app.NewConfigurationManager()
	dbPool := postgresql.GetConnectionPool(ctx, configurationManager.PostgresqlConfig)
	productRepository := repository.NewProductRepository(dbPool)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	e := echo.New()
	productController.RegisterRoutes(e)
	if err := e.Start("localhost:8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
