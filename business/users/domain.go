package users

import (
	"time"
)

type Domain struct {
	ID           int
	Username     string
	Email        string
	Password     string
	Name         string
	Dob          string
	Phone_Number string
	Address      string
	City         string
	Province     string
	Post_code    int
	Photo        string
	Token        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Service interface {
	Register(domain *Domain) (Domain, error)
	Login(email, password string) (Domain, error)
	Update(userID int, domain *Domain) (Domain, error)
	UserByID(id int) (Domain, error)
}

type Repository interface {
	Register(domain *Domain) (Domain, error)
	Login(email, password string) (Domain, error)
	Update(userID int, domain *Domain) (Domain, error)
	UserByID(id int) (Domain, error)
}
