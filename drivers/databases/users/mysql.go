package users

import (
	"rulzmotoshop/business"
	"rulzmotoshop/business/users"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) users.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (rep *MysqlUserRepository) Register(domain *users.Domain) (users.Domain, error) {

	user := fromDomain(*domain)

	result := rep.Conn.Create(&user)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return toDomain(user), nil
}

func (rep *MysqlUserRepository) Login(email, password string) (users.Domain, error) {
	var user Users
	err := rep.Conn.First(&user, "email = ?", email).Error

	if err != nil {
		return users.Domain{}, business.ErrEmailorPass
	}

	return toDomain(user), nil
}
func (rep *MysqlUserRepository) Update(userID int, domain *users.Domain) (users.Domain, error) {

	profileUpdate := fromDomain(*domain)

	profileUpdate.ID = userID

	result := rep.Conn.Where("id = ?", userID).Updates(&profileUpdate)

	if result.Error != nil {
		return users.Domain{}, business.ErrNotFound
	}

	return toDomainUpdate(profileUpdate), nil
}
func (rep *MysqlUserRepository) UserByID(id int) (users.Domain, error) {

	var user Users

	result := rep.Conn.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return toDomain(user), nil
}
