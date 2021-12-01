package drivers

import (
	adminDomain "rulzmotoshop/business/admins"
	itemDomain "rulzmotoshop/business/items"
	sellerDomain "rulzmotoshop/business/sellers"
	wishDomain "rulzmotoshop/business/wishlist"

	transDomain "rulzmotoshop/business/transactions"
	userDomain "rulzmotoshop/business/users"

	adminDB "rulzmotoshop/drivers/databases/admins"
	itemDB "rulzmotoshop/drivers/databases/items"
	sellerDB "rulzmotoshop/drivers/databases/sellers"

	transDB "rulzmotoshop/drivers/databases/transactions"
	userDB "rulzmotoshop/drivers/databases/users"
	wishDB "rulzmotoshop/drivers/databases/wishlist"

	"gorm.io/gorm"
)

func NewAdminRepository(conn *gorm.DB) adminDomain.Repository {
	return adminDB.NewMysqlAdminRepository(conn)
}
func NewSellerRepository(conn *gorm.DB) sellerDomain.Repository {
	return sellerDB.NewMysqlSellerRepository(conn)
}
func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMysqlUserRepository(conn)
}
func NewItemRepository(conn *gorm.DB) itemDomain.Repository {
	return itemDB.NewMysqlItemRepository(conn)
}
func NewTransRepository(conn *gorm.DB) transDomain.Repository {
	return transDB.NewMysqlTransRepository(conn)
}
func NewWishRepository(conn *gorm.DB) wishDomain.Repository {
	return wishDB.NewMysqlWishRepository(conn)
}
