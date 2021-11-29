package admins

import (
	"net/http"
	"rulzmotoshop/business/admins"
	"rulzmotoshop/controllers"
	"rulzmotoshop/controllers/admins/request"
	"rulzmotoshop/controllers/admins/response"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	adminService admins.Service
}

func NewControllerAdmin(serv admins.Service) *AdminController {
	return &AdminController{
		adminService: serv,
	}
}

func (ctrl *AdminController) Register(c echo.Context) error {

	registerReq := request.Admins{}

	if err := c.Bind(&registerReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	result, err := ctrl.adminService.Register(registerReq.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainRegister(result))

}

func (ctrl *AdminController) Login(c echo.Context) error {

	loginReq := request.AdminLogin{}

	if err := c.Bind(&loginReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	result, err := ctrl.adminService.Login(loginReq.Username, loginReq.Password)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainLogin(result))
}
