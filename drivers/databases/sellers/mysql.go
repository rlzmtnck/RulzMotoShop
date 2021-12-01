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
func (rep *MysqlSellersRepository) Update(sellID int, domain *sellers.Domain) (sellers.Domain, error) {

	profileUpdate := fromDomain(*domain)

	profileUpdate.ID = sellID

	result := rep.Conn.Where("id = ?", sellID).Updates(&profileUpdate)

	if result.Error != nil {
		return sellers.Domain{}, business.ErrNotFound
	}

	return toDomainUpdate(profileUpdate), nil
}
func (rep *MysqlSellersRepository) SellerByID(id int) (sellers.Domain, error) {

	var seller Sellers

	result := rep.Conn.Where("id = ?", id).First(&seller)

	if result.Error != nil {
		return sellers.Domain{}, result.Error
	}

	return toDomain(seller), nil
}
func (rep *MysqlSellersRepository) Delete(id int) (string, error) {
	rec := Sellers{}

	find := rep.Conn.Where("id = ?", id).First(&rec)

	if find.Error != nil {
		return "", business.ErrUnathorized
	}

	err := rep.Conn.Delete(&rec, "id = ?", id).Error

	if err != nil {
		return "", business.ErrNotFound
	}

	return "Seller has been delete", nil

}
