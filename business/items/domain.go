package items

import "time"

type Domain struct {
	ID          int
	SellerID    int
	Name        string
	Description string
	Color       string
	Stock       int
	Poster      string
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Service interface {
	AllItem() ([]Domain, error)
	Create(sellID int, domain *Domain) (Domain, error)
	Update(sellID int, itID int, domain *Domain) (Domain, error)
	Delete(sellID, id int) (string, error)
	DeleteByAdmin(id int) (string, error)
	MyItemBySeller(sellID int) ([]Domain, error)
	ItemByID(id int) (Domain, error)
	ItemByIdSeller(sellsID int) ([]Domain, error)
}

type Repository interface {
	AllItem() ([]Domain, error)
	Create(sellID int, domain *Domain) (Domain, error)
	Update(sellID int, itID int, domain *Domain) (Domain, error)
	Delete(sellID, id int) (string, error)
	DeleteByAdmin(id int) (string, error)
	MyItemBySeller(sellID int) ([]Domain, error)
	ItemByID(id int) (Domain, error)
	ItemByIdSeller(sellsID int) ([]Domain, error)
}
