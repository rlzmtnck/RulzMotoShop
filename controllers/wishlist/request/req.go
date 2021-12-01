package request

import (
	"rulzmotoshop/business/wishlist"
)

type Wishlist struct {
	ItemID int `json:"item_id"`
}

func (req *Wishlist) ToDomain() *wishlist.Domain {
	return &wishlist.Domain{
		ItemID: req.ItemID,
	}
}
