package wishlist

import "time"

type Domain struct {
	ID        int
	UserID    int
	ItemID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Wish(userID int, domain *Domain) (Domain, error)
	GetAllUserWish(userID int) ([]Domain, error)
}

type Repository interface {
	Wish(userID int, domain *Domain) (Domain, error)
	GetAllUserWish(userID int) ([]Domain, error)
}
