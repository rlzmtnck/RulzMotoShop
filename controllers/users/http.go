package users

import (
	"net/http"
	"rulzmotoshop/business/users"
	"rulzmotoshop/controllers"
	"rulzmotoshop/controllers/users/request"
	"rulzmotoshop/controllers/users/response"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService users.Service
}

func NewControllerUser(serv users.Service) *UserController {
	return &UserController{
		userService: serv,
	}
}

func (ctrl *UserController) Register(c echo.Context) error {

	// ctx := c.Request().Context()
	registerReq := request.Users{}

	if err := c.Bind(&registerReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	result, err := ctrl.userService.Register(registerReq.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainRegister(result))

}

func (ctrl *UserController) Login(c echo.Context) error {

	loginReq := request.UserLogin{}

	if err := c.Bind(&loginReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	result, err := ctrl.userService.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainLogin(result))
}
