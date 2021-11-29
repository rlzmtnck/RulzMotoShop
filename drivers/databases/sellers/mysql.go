package sellers

import (
	"rulzmotoshop/business"
	"rulzmotoshop/business/sellers"

	"gorm.io/gorm"
)

type MysqlSellersRepository struct {
	Conn *gorm.DB
}

func NewMysqlSellerRepository(conn *gorm.DB) sellers.Repository {
	return &MysqlSellersRepository{
		Conn: conn,
	}
}

func (rep *MysqlSellersRepository) Register(domain *sellers.Domain) (sellers.Domain, error) {

	org := fromDomain(*domain)

	result := rep.Conn.Create(&org)

	if result.Error != nil {
		return sellers.Domain{}, result.Error
	}

	return toDomain(org), nil
}

func (rep *MysqlSellersRepository) Login(email, password string) (sellers.Domain, error) {
	var org Sellers
	err := rep.Conn.First(&org, "email = ?", email).Error

	if err != nil {
		return sellers.Domain{}, business.ErrEmailorPass
	}

	return toDomain(org), nil
}
