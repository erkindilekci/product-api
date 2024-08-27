package request

import "github.com/erkindilekci/product-app/pkg/service/dto"

type AddProductRequest struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
	Store    string  `json:"store"`
}

func (request *AddProductRequest) ToModel() dto.ProductCreate {
	return dto.ProductCreate{
		Name:     request.Name,
		Price:    request.Price,
		Discount: request.Discount,
		Store:    request.Store,
	}
}
