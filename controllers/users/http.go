package users

import (
	"net/http"
	"rulzmotoshop/business/users"
	"rulzmotoshop/controllers"
	"rulzmotoshop/controllers/users/request"
	"rulzmotoshop/controllers/users/response"
	"strconv"

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
func (ctrl *UserController) Update(c echo.Context) error {

	updateReq := request.Users{}

	if err := c.Bind(&updateReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))

	getData, _ := ctrl.userService.UserByID(id)
	result, err := ctrl.userService.Update(id, updateReq.ToDomain())
	result.ID = getData.ID

	result.Name = getData.Name

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainUpdateUser(result))

}
func (ctrl *UserController) UserByID(c echo.Context) error {

	userID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.userService.UserByID(userID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, response.FromDomainAllUser(result))
}
