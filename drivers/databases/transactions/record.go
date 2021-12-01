package transactions

import (
	"rulzmotoshop/business/transactions"
	"time"

	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model
	ID         int `gorm:"primary_key"`
	UserID     int
	ItemID     int
	Status     bool
	Trans_code string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func toDomain(tr Transactions) transactions.Domain {
	return transactions.Domain{
		ID:         tr.ID,
		UserID:     tr.UserID,
		ItemID:     tr.ItemID,
		Status:     tr.Status,
		Trans_code: tr.Trans_code,
		CreatedAt:  tr.CreatedAt,
		UpdatedAt:  tr.UpdatedAt,
	}
}

func fromDomain(domain transactions.Domain) Transactions {
	return Transactions{
		ID:         domain.ID,
		UserID:     domain.UserID,
		ItemID:     domain.ItemID,
		Status:     domain.Status,
		Trans_code: domain.Trans_code,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}

func toDomainList(data []Transactions) []transactions.Domain {
	result := []transactions.Domain{}

	for _, trans := range data {
		result = append(result, toDomain(trans))
	}
	return result
}
