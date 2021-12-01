package main

import (
	"log"

	_routes "rulzmotoshop/app/routes"

	_userService "rulzmotoshop/business/users"
	_userController "rulzmotoshop/controllers/users"
	_userRepo "rulzmotoshop/drivers/databases/users"

	_wishService "rulzmotoshop/business/wishlist"
	_wishController "rulzmotoshop/controllers/wishlist"
	_wishRepo "rulzmotoshop/drivers/databases/wishlist"

	_sellerService "rulzmotoshop/business/sellers"
	_sellerController "rulzmotoshop/controllers/sellers"
	_sellerRepo "rulzmotoshop/drivers/databases/sellers"

	_itemService "rulzmotoshop/business/items"
	_itemController "rulzmotoshop/controllers/items"
	_itemsRepo "rulzmotoshop/drivers/databases/items"

	_transService "rulzmotoshop/business/transactions"
	_transController "rulzmotoshop/controllers/transactions"
	_transRepo "rulzmotoshop/drivers/databases/transactions"

	_adminService "rulzmotoshop/business/admins"
	_adminController "rulzmotoshop/controllers/admins"
	_adminRepo "rulzmotoshop/drivers/databases/admins"

	_newsController "rulzmotoshop/controllers/news"
	_newsRepo "rulzmotoshop/drivers/databases/thirdparty/news"

	_dbDriver "rulzmotoshop/drivers/mysql"

	_driverFactory "rulzmotoshop/drivers"

	_middleware "rulzmotoshop/app/middleware"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_userRepo.Users{},
		&_sellerRepo.Sellers{},
		&_itemsRepo.Items{},
		&_transRepo.Transactions{},
		&_adminRepo.Admins{},
		&_wishRepo.Wishlist{},
	)
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitDB()
	dbMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: int64(viper.GetInt(`jwt.expired`)),
	}

	e := echo.New()

	userRepo := _driverFactory.NewUserRepository(db)
	userService := _userService.NewServiceUser(userRepo, 10, &configJWT)
	userCtrl := _userController.NewControllerUser(userService)

	wishRepo := _driverFactory.NewWishRepository(db)
	wishService := _wishService.NewServiceWish(wishRepo)
	wishCtrl := _wishController.NewControllerWish(wishService)

	SellerRepo := _driverFactory.NewSellerRepository(db)
	sellerrService := _sellerService.NewServiceSeller(SellerRepo, 10, &configJWT)
	sellerCtrl := _sellerController.NewControllerSeller(sellerrService)

	itemRepo := _driverFactory.NewItemRepository(db)
	itemService := _itemService.NewServiceItem(itemRepo)
	itemCtrl := _itemController.NewControllerItem(itemService)

	transRepo := _driverFactory.NewTransRepository(db)
	transService := _transService.NewServiceTrans(transRepo)
	transCtrl := _transController.NewControllerItem(transService)

	adminRepo := _driverFactory.NewAdminRepository(db)
	adminService := _adminService.NewServiceAdmin(adminRepo, 10, &configJWT)
	adminCtrl := _adminController.NewControllerAdmin(adminService)

	newsRepo := _newsRepo.NewNewsApi()
	newsCtrl := _newsController.NewNewsController(newsRepo)

	routesInit := _routes.ControllerList{
		JWTMiddleware:    configJWT.Init(),
		UserController:   *userCtrl,
		SellerController: *sellerCtrl,
		ItemController:   *itemCtrl,
		TransController:  *transCtrl,
		AdminController:  *adminCtrl,
		WishController:   *wishCtrl,
		NewsController:   *newsCtrl,
	}

	routesInit.RouteRegister(e)
	_middleware.Logger(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
