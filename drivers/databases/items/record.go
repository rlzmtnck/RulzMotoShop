package items

import (
	"rulzmotoshop/business/items"
	"time"

	"gorm.io/gorm"
)

type Items struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	SellerID    int
	Name        string
	Description string
	Color       string
	Stock       int
	Poster      string
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func toDomain(ev Items) items.Domain {
	return items.Domain{
		ID:          ev.ID,
		SellerID:    ev.SellerID,
		Name:        ev.Name,
		Description: ev.Description,
		Color:       ev.Color,
		Stock:       ev.Stock,
		Poster:      ev.Poster,
		Price:       ev.Price,
		CreatedAt:   ev.CreatedAt,
		UpdatedAt:   ev.UpdatedAt,
	}
}

func fromDomain(domain items.Domain) Items {
	return Items{
		ID:          domain.ID,
		SellerID:    domain.SellerID,
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

func toDomainUpdate(ev Items) items.Domain {
	return items.Domain{
		ID:          ev.ID,
		SellerID:    ev.SellerID,
		Name:        ev.Name,
		Description: ev.Description,
		Color:       ev.Color,
		Stock:       ev.Stock,
		Poster:      ev.Poster,
		Price:       ev.Price,
		CreatedAt:   ev.CreatedAt,
		UpdatedAt:   ev.UpdatedAt,
	}
}
func toDomainList(data []Items) []items.Domain {
	result := []items.Domain{}

	for _, ev := range data {
		result = append(result, toDomain(ev))
	}
	return result
}
