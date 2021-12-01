package wishlist

import (
	"rulzmotoshop/business/wishlist"
	"time"

	"gorm.io/gorm"
)

type Wishlist struct {
	gorm.Model
	ID        int `gorm:"primary_key"`
	UserID    int
	ItemID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toDomain(ws Wishlist) wishlist.Domain {
	return wishlist.Domain{
		ID:        ws.ID,
		UserID:    ws.UserID,
		ItemID:    ws.ItemID,
		CreatedAt: ws.CreatedAt,
		UpdatedAt: ws.UpdatedAt,
	}
}

func fromDomain(domain wishlist.Domain) Wishlist {
	return Wishlist{
		ID:        domain.ID,
		UserID:    domain.UserID,
		ItemID:    domain.ItemID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func toDomainList(data []Wishlist) []wishlist.Domain {
	result := []wishlist.Domain{}

	for _, wish := range data {
		result = append(result, toDomain(wish))
	}
	return result
}
