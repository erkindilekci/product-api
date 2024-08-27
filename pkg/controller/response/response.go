package response

import "github.com/erkindilekci/product-app/pkg/domain"

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

func NewErrorResponse(errorMessage string) *ErrorResponse {
	return &ErrorResponse{errorMessage}
}

type ProductResponse struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
	Store    string  `json:"store"`
}

func ToProductResponse(product domain.Product) ProductResponse {
	return ProductResponse{
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Price,
		Store:    product.Store,
	}
}

func ToProductResponseList(products []domain.Product) []ProductResponse {
	var responses []ProductResponse
	for _, product := range products {
		responses = append(responses, ToProductResponse(product))
	}
	return responses
}
