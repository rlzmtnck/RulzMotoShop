package sellers

import (
	"rulzmotoshop/business/sellers"
	"time"

	"gorm.io/gorm"
)

type Sellers struct {
	gorm.Model
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"unique"`
	Email        string `gorm:"unique"`
	Password     string
	Name         string
	Shop_Name    string
	Phone_Number string
	Photo        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func toDomain(org Sellers) sellers.Domain {
	return sellers.Domain{
		ID:           org.ID,
		Username:     org.Username,
		Email:        org.Email,
		Password:     org.Password,
		Name:         org.Name,
		Shop_Name:    org.Shop_Name,
		Phone_Number: org.Phone_Number,
		Photo:        org.Photo,
		CreatedAt:    org.CreatedAt,
		UpdatedAt:    org.UpdatedAt,
	}
}
func toDomainUpdate(upd Sellers) sellers.Domain {
	return sellers.Domain{
		ID:           upd.ID,
		Username:     upd.Username,
		Email:        upd.Email,
		Password:     upd.Password,
		Name:         upd.Name,
		Shop_Name:    upd.Shop_Name,
		Phone_Number: upd.Phone_Number,
		Photo:        upd.Photo,
		CreatedAt:    upd.CreatedAt,
		UpdatedAt:    upd.UpdatedAt,
	}
}
func fromDomain(domain sellers.Domain) Sellers {
	return Sellers{
		ID:           domain.ID,
		Username:     domain.Username,
		Email:        domain.Email,
		Password:     domain.Password,
		Name:         domain.Name,
		Shop_Name:    domain.Shop_Name,
		Phone_Number: domain.Phone_Number,
		Photo:        domain.Photo,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
