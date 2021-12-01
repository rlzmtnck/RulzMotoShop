package response

import (
	"rulzmotoshop/business/sellers"
	"time"
)

type SellerRegisterRespons struct {
	Message   string    `json:"message"`
	ID        int       `json:"id:"`
	Username  string    `json:"username"`
	Shop_Name string    `json:"shop_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type SellerResponse struct {
	ID        int       `json:"id:"`
	Username  string    `json:"username"`
	Shop_Name string    `json:"shop_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomainRegister(domain sellers.Domain) SellerRegisterRespons {
	return SellerRegisterRespons{
		Message:   "Seller Registration Success",
		ID:        domain.ID,
		Shop_Name: domain.Shop_Name,
		Username:  domain.Username,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
func FromDomainAllSeller(domain sellers.Domain) SellerResponse {
	return SellerResponse{
		ID:        domain.ID,
		Shop_Name: domain.Shop_Name,
		Username:  domain.Username,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

type SellerLoginRespons struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func FromDomainLogin(domain sellers.Domain) SellerLoginRespons {
	return SellerLoginRespons{
		Message: "Seller Login Success",
		Token:   domain.Token,
	}
}
func FromDomainUpdateSeller(domain sellers.Domain) SellerRegisterRespons {
	return SellerRegisterRespons{
		Message:   "Update Shop Success",
		ID:        domain.ID,
		Shop_Name: domain.Shop_Name,
		Username:  domain.Username,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
