package request

import (
	"rulzmotoshop/business/sellers"
)

type Sellers struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Shop_Name    string `json:"shop_name"`
	Phone_Number string `json:"phone_number"`
	Photo        string `json:"photo"`
}
type SellerLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *Sellers) ToDomain() *sellers.Domain {
	return &sellers.Domain{
		Username:     req.Username,
		Password:     req.Password,
		Email:        req.Email,
		Name:         req.Name,
		Shop_Name:    req.Shop_Name,
		Phone_Number: req.Phone_Number,
		Photo:        req.Photo,
	}
}
