package sellers

import (
	"net/http"
	"rulzmotoshop/business/sellers"
	"rulzmotoshop/controllers"
	"rulzmotoshop/controllers/sellers/request"
	"rulzmotoshop/controllers/sellers/response"

	"github.com/labstack/echo/v4"
)

type SellerController struct {
	sellerService sellers.Service
}

func NewControllerSeller(serv sellers.Service) *SellerController {
	return &SellerController{
		sellerService: serv,
	}
}

func (ctrl *SellerController) Register(c echo.Context) error {

	registerReq := request.Sellers{}

	if err := c.Bind(&registerReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	result, err := ctrl.sellerService.Register(registerReq.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainRegister(result))

}

func (ctrl *SellerController) Login(c echo.Context) error {

	loginReq := request.SellerLogin{}

	if err := c.Bind(&loginReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	result, err := ctrl.sellerService.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainLogin(result))
}
