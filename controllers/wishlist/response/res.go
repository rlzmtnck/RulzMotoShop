package response

import (
	"rulzmotoshop/business/wishlist"
)

type CreateWishResponse struct {
	Message string `json:"message"`
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	ItemID  int    `json:"item_id"`
}

func FromDomainCreate(domain wishlist.Domain) CreateWishResponse {
	return CreateWishResponse{
		Message: "Success, Wishlist added to this item",
		ID:      domain.ID,
		UserID:  domain.UserID,
		ItemID:  domain.ItemID,
	}
}

type WishResponse struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	ItemID int `json:"item_id"`
}

func FromDomainAllWish(domain wishlist.Domain) WishResponse {
	return WishResponse{
		ID:     domain.ID,
		UserID: domain.UserID,
		ItemID: domain.ItemID,
	}
}

func FromWishListDomain(domain []wishlist.Domain) []WishResponse {
	var response []WishResponse
	for _, value := range domain {
		response = append(response, FromDomainAllWish(value))
	}
	return response
}
