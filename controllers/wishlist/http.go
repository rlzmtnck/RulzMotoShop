package wishlist

import (
	"net/http"

	"rulzmotoshop/app/middleware"
	"rulzmotoshop/business/wishlist"
	"rulzmotoshop/controllers"
	"rulzmotoshop/controllers/wishlist/request"
	"rulzmotoshop/controllers/wishlist/response"

	"github.com/labstack/echo/v4"
)

type WishController struct {
	wishService wishlist.Service
}

func NewControllerWish(serv wishlist.Service) *WishController {
	return &WishController{
		wishService: serv,
	}
}

func (ctrl *WishController) Create(c echo.Context) error {

	createReq := request.Wishlist{}

	if err := c.Bind(&createReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	jwtGetID := middleware.GetUser(c)

	result, err := ctrl.wishService.Wish(jwtGetID.ID, createReq.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainCreate(result))

}

func (ctrl *WishController) GetAllUserWish(c echo.Context) error {

	jwtGetID := middleware.GetUser(c)

	result, err := ctrl.wishService.GetAllUserWish(jwtGetID.ID)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromWishListDomain(result))
}
