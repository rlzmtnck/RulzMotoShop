package request

import (
	"rulzmotoshop/business/transactions"
)

type Transactions struct {
	ItemID int `json:"item_id"`
}

func (req *Transactions) ToDomain() *transactions.Domain {
	return &transactions.Domain{
		ItemID: req.ItemID,
	}
}
