package users

import (
	"rulzmotoshop/business/users"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID           int       `json:"id" form:"id" gorm:"primary_key"`
	Username     string    `json:"username" form:"username" gorm:"unique"`
	Email        string    `json:"email" form:"email" gorm:"unique"`
	Password     string    `json:"password" form:"password"`
	Name         string    `json:"name" form:"name"`
	Dob          string    `json:"dob" form:"dob"`
	Phone_Number string    `json:"phone_number" form:"phone_number"`
	Photo        string    `json:"photo" form:"photo"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" form:"updated_at"`
}

func toDomain(user Users) users.Domain {
	return users.Domain{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		Password:     user.Password,
		Name:         user.Name,
		Dob:          user.Dob,
		Phone_Number: user.Phone_Number,
		Photo:        user.Photo,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func fromDomain(domain users.Domain) Users {
	return Users{
		ID:           domain.ID,
		Username:     domain.Username,
		Email:        domain.Email,
		Password:     domain.Password,
		Name:         domain.Name,
		Dob:          domain.Dob,
		Phone_Number: domain.Phone_Number,
		Photo:        domain.Photo,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
