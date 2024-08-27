package controller

import (
	"fmt"
	"github.com/erkindilekci/product-app/pkg/controller/request"
	"github.com/erkindilekci/product-app/pkg/controller/response"
	"github.com/erkindilekci/product-app/pkg/domain"
	"github.com/erkindilekci/product-app/pkg/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ProductController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{productService}
}

func (controller *ProductController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/products", controller.GetAllProducts)
	e.GET("/api/v1/products/:id", controller.GetProductById)
	e.POST("/api/v1/products", controller.AddNewProduct)
	e.PUT("/api/v1/products/:id", controller.UpdatePriceById)
	e.DELETE("/api/v1/products/:id", controller.DeleteProductById)
}

func (controller *ProductController) GetAllProducts(c echo.Context) error {
	store := c.QueryParam("store")
	var products []domain.Product

	if len(store) == 0 {
		products = controller.productService.GetAllProducts()
	} else {
		products = controller.productService.GetProductsByStore(store)
	}

	return c.JSON(http.StatusOK, response.ToProductResponseList(products))
}

func (controller *ProductController) GetProductById(c echo.Context) error {
	param := c.Param("id")
	if param == "" {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request: no product id specified"))
	}

	productId, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request: product id must be an integer"))
	}

	product, err := controller.productService.GetById(int64(productId))
	if err != nil {
		return c.JSON(http.StatusNotFound, response.NewErrorResponse(fmt.Sprintf("Product not found: no product with ID %d", productId)))
	}

	return c.JSON(http.StatusOK, response.ToProductResponse(product))
}

func (controller *ProductController) AddNewProduct(c echo.Context) error {
	var addProductRequest request.AddProductRequest
	err := c.Bind(&addProductRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request: unable to bind the provided data to the product structure"))
	}

	err = controller.productService.Add(addProductRequest.ToModel())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.NewErrorResponse(err.Error()))
	}

	return c.NoContent(http.StatusCreated)
}

func (controller *ProductController) UpdatePriceById(c echo.Context) error {
	param := c.Param("id")
	if param == "" {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request: no product id specified"))
	}

	productId, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request: product id must be an integer"))
	}

	newPrice := c.QueryParam("newPrice")
	if len(newPrice) == 0 {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request: no newPrice query parameter found"))
	}

	priceFloat, err := strconv.ParseFloat(newPrice, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request: newPrice must be a float"))
	}

	err = controller.productService.UpdatePrice(int64(productId), float32(priceFloat))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse(err.Error()))
	}

	return c.NoContent(http.StatusOK)
}

func (controller *ProductController) DeleteProductById(c echo.Context) error {
	param := c.Param("id")
	if param == "" {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request: no product id specified"))
	}

	productId, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request: product id must be an integer"))
	}

	err = controller.productService.DeleteById(int64(productId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse(err.Error()))
	}

	return c.NoContent(http.StatusOK)
}
