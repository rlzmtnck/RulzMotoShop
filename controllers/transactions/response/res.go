package response

import "rulzmotoshop/business/transactions"

type CreateTransResponse struct {
	Message    string `json:"message"`
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	ItemID     int    `json:"item_id"`
	Status     bool   `json:"status"`
	Trans_code string `json:"trans_code"`
}

func FromDomainCreate(domain transactions.Domain) CreateTransResponse {
	return CreateTransResponse{
		Message:    "Transactions Success",
		ID:         domain.ID,
		UserID:     domain.UserID,
		ItemID:     domain.ItemID,
		Status:     domain.Status,
		Trans_code: domain.Trans_code,
	}
}

type TransResponse struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	ItemID     int    `json:"item_id"`
	Status     bool   `json:"status"`
	Trans_code string `json:"trans_code"`
}

func FromDomainAllTrans(domain transactions.Domain) TransResponse {
	return TransResponse{
		ID:         domain.ID,
		UserID:     domain.UserID,
		ItemID:     domain.ItemID,
		Status:     domain.Status,
		Trans_code: domain.Trans_code,
	}
}

func FromTransListDomain(domain []transactions.Domain) []TransResponse {
	var response []TransResponse
	for _, value := range domain {
		response = append(response, FromDomainAllTrans(value))
	}
	return response
}
