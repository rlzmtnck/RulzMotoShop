package request

import "rulzmotoshop/business/items"

type Items struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
	Color       string `json:"color"`
	Stock       int    `json:"stock"`
	Poster      string `json:"poster"`
	Price       int    `json:"price"`
}

func (req *Items) ToDomain() *items.Domain {
	return &items.Domain{
		Name:        req.Name,
		Description: req.Description,
		Color:       req.Color,
		Stock:       req.Stock,
		Poster:      req.Poster,
		Price:       req.Price,
	}
}
