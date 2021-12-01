package sellers

import "time"

type Domain struct {
	ID           int
	Username     string
	Email        string
	Password     string
	Name         string
	Shop_Name    string
	Phone_Number string
	Photo        string
	Token        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Service interface {
	Register(domain *Domain) (Domain, error)
	Login(email, password string) (Domain, error)
	Update(sellID int, domain *Domain) (Domain, error)
	SellerByID(id int) (Domain, error)
	Delete(id int) (string, error)
}

type Repository interface {
	Register(domain *Domain) (Domain, error)
	Login(email, password string) (Domain, error)
	Update(sellID int, domain *Domain) (Domain, error)
	SellerByID(id int) (Domain, error)
	Delete(id int) (string, error)
}
