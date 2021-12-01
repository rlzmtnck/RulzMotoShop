package items

import (
	"net/http"
	"rulzmotoshop/app/middleware"
	"rulzmotoshop/business/items"
	"rulzmotoshop/controllers"
	"rulzmotoshop/controllers/items/request"
	"rulzmotoshop/controllers/items/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ItemController struct {
	itemService items.Service
}

func NewControllerItem(serv items.Service) *ItemController {
	return &ItemController{
		itemService: serv,
	}
}

func (ctrl *ItemController) Create(c echo.Context) error {

	createReq := request.Items{}

	if err := c.Bind(&createReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	jwtGetID := middleware.GetUser(c)

	result, err := ctrl.itemService.Create(jwtGetID.ID, createReq.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainCreate(result))

}

func (ctrl *ItemController) AllItem(c echo.Context) error {

	result, err := ctrl.itemService.AllItem()

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromItemListDomain(result))

}

func (ctrl *ItemController) Update(c echo.Context) error {

	updateReq := request.Items{}

	if err := c.Bind(&updateReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	jwtGetID := middleware.GetUser(c)
	getData, _ := ctrl.itemService.ItemByID(id)
	result, err := ctrl.itemService.Update(jwtGetID.ID, id, updateReq.ToDomain())
	result.ID = getData.ID
	result.Name = getData.Name
	result.Description = getData.Description

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainUpdateItem(result))

}

func (ctrl *ItemController) Delete(c echo.Context) error {

	orgzID := middleware.GetUser(c)
	deletedId, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.itemService.Delete(orgzID.ID, deletedId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, result)

}
func (ctrl *ItemController) DeleteByAdmin(c echo.Context) error {

	deletedId, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.itemService.DeleteByAdmin(deletedId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, result)

}
func (ctrl *ItemController) MyItemBySeller(c echo.Context) error {
	orgzID := middleware.GetUser(c)

	result, err := ctrl.itemService.MyItemBySeller(orgzID.ID)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromItemListDomain(result))
}

func (ctrl *ItemController) ItemByID(c echo.Context) error {

	itemID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.itemService.ItemByID(itemID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, response.FromDomainAllItem(result))
}

func (ctrl *ItemController) ItemByIdSeller(c echo.Context) error {

	sellsID, _ := strconv.Atoi(c.Param("sellerID"))

	result, err := ctrl.itemService.ItemByIdSeller(sellsID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromItemListDomain(result))
}
