package wishlist

import (
	"rulzmotoshop/business/wishlist"

	"gorm.io/gorm"
)

type MysqlWishRepository struct {
	Conn *gorm.DB
}

func NewMysqlWishRepository(conn *gorm.DB) wishlist.Repository {
	return &MysqlWishRepository{
		Conn: conn,
	}
}

func (rep *MysqlWishRepository) Wish(userID int, domain *wishlist.Domain) (wishlist.Domain, error) {

	tr := fromDomain(*domain)

	tr.UserID = userID

	result := rep.Conn.Create(&tr)

	if result.Error != nil {
		return wishlist.Domain{}, result.Error
	}

	return toDomain(tr), nil
}

func (rep *MysqlWishRepository) GetAllUserWish(userID int) ([]wishlist.Domain, error) {

	wish := []Wishlist{}

	result := rep.Conn.Where("user_id = ?", userID).Find(&wish)

	if result.Error != nil {
		return []wishlist.Domain{}, result.Error
	}

	return toDomainList(wish), nil
}
