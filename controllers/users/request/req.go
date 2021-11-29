package request

import "rulzmotoshop/business/users"

type Users struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Dob          string `json:"dob"`
	Phone_Number string `json:"phone_number"`
	Photo        string `json:"photo"`
}

func (req *Users) ToDomain() *users.Domain {
	return &users.Domain{
		Username:     req.Username,
		Password:     req.Password,
		Email:        req.Email,
		Name:         req.Name,
		Dob:          req.Dob,
		Phone_Number: req.Phone_Number,
		Photo:        req.Photo,
	}
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
