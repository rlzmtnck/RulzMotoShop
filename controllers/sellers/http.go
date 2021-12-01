package sellers

import (
	"net/http"

	"rulzmotoshop/business/sellers"
	"rulzmotoshop/controllers"
	"rulzmotoshop/controllers/sellers/request"
	"rulzmotoshop/controllers/sellers/response"
	"strconv"

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
func (ctrl *SellerController) Update(c echo.Context) error {

	updateReq := request.Sellers{}

	if err := c.Bind(&updateReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))

	getData, _ := ctrl.sellerService.SellerByID(id)
	result, err := ctrl.sellerService.Update(id, updateReq.ToDomain())
	result.ID = getData.ID

	result.Name = getData.Name

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainUpdateSeller(result))

}
func (ctrl *SellerController) SellerByID(c echo.Context) error {

	itemID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.sellerService.SellerByID(itemID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, response.FromDomainAllSeller(result))
}
func (ctrl *SellerController) Delete(c echo.Context) error {

	deletedId, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.sellerService.Delete(deletedId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, result)

}
