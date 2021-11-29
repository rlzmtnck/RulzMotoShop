package routes

import (
	"net/http"
	middlewareApp "rulzmotoshop/app/middleware"
	"rulzmotoshop/business"
	controller "rulzmotoshop/controllers"
	"rulzmotoshop/controllers/admins"
	"rulzmotoshop/controllers/items"
	"rulzmotoshop/controllers/sellers"
	"rulzmotoshop/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware    middleware.JWTConfig
	UserController   users.UserController
	SellerController sellers.SellerController
	ItemController   items.ItemController
	// TransController     transactions.TransController
	AdminController admins.AdminController
	// NewsController      news.NewsController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	// Public
	e.GET("/items", cl.ItemController.AllItem)
	e.GET("/item/:id", cl.ItemController.ItemByID)
	e.GET("/:sellerID/items", cl.ItemController.ItemByIdSeller)
	// e.GET("/news", cl.NewsController.GetNewsByCategory)

	// Admins
	admins := e.Group("admins")
	admins.POST("/register", cl.AdminController.Register)
	admins.POST("/login", cl.AdminController.Login)
	admins.PUT("/update-seller/:id", cl.SellerController.Update, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())

	// Users
	users := e.Group("users")
	users.POST("/register", cl.UserController.Register)
	users.POST("/login", cl.UserController.Login)
	// users.GET("/my-events", cl.TransController.GetAllUserTrans, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationUser())

	// Seller
	sellers := e.Group("sellers")
	sellers.POST("/register", cl.SellerController.Register)
	sellers.POST("/login", cl.SellerController.Login)
	sellers.PUT("/update-shop/:id", cl.SellerController.Update, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationSeller())

	sellers.POST("/add-item", cl.ItemController.Create, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationSeller())
	sellers.PUT("/update-item/:id", cl.ItemController.Update, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationSeller())

	sellers.DELETE("/delete-item/:id", cl.ItemController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationSeller())
	sellers.GET("/my-item", cl.ItemController.MyItemBySeller, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationSeller())

	// // Transaction
	// e.POST("/transaction", cl.TransController.Create, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationUser())
	// e.GET("/all-transactions", cl.TransController.GetAllTrans, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	// e.GET("/transaction/:id", cl.TransController.GetTransByID, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())

}

func RoleValidationAdmin() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			if claims.Role == "admin" {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, business.ErrUnathorized)
			}
		}
	}
}

func RoleValidationUser() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			if claims.Role == "user" || claims.Role == "admin" {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, business.ErrUnathorized)
			}
		}
	}
}

func RoleValidationSeller() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			if claims.Role == "seller" || claims.Role == "admin" {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, business.ErrUnathorized)
			}
		}
	}
}
