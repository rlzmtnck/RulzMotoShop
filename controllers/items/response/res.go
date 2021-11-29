package response

import (
	"rulzmotoshop/business/items"
	"time"
)

type CreateItemResponse struct {
	Message     string    `json:"message"`
	ID          int       `json:"id:"`
	Name        string    `json:"name"`
	Description string    `json:"desc"`
	Color       string    `json:"color"`
	Stock       int       `json:"stock"`
	Poster      string    `json:"poster"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomainCreate(domain items.Domain) CreateItemResponse {
	return CreateItemResponse{
		Message:     "Add Item Success",
		ID:          domain.ID,
		Name:        domain.Name,
		Description: domain.Description,
		Color:       domain.Color,
		Stock:       domain.Stock,
		Poster:      domain.Poster,
		Price:       domain.Price,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

type ItemResponse struct {
	ID          int       `json:"id:"`
	Name        string    `json:"name"`
	Description string    `json:"desc"`
	Color       string    `json:"color"`
	Stock       int       `json:"stock"`
	Poster      string    `json:"poster"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomainAllItem(domain items.Domain) ItemResponse {
	return ItemResponse{
		ID:          domain.ID,
		Name:        domain.Name,
		Description: domain.Description,
		Color:       domain.Color,
		Stock:       domain.Stock,
		Poster:      domain.Poster,
		Price:       domain.Price,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func FromDomainUpdateItem(domain items.Domain) CreateItemResponse {
	return CreateItemResponse{
		Message:   "Update Item Success",
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromItemListDomain(domain []items.Domain) []ItemResponse {
	var response []ItemResponse
	for _, value := range domain {
		response = append(response, FromDomainAllItem(value))
	}
	return response
}
