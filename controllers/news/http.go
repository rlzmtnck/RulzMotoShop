package news

import (
	"fmt"
	"net/http"
	"rulzmotoshop/business/news"
	"rulzmotoshop/controllers"

	"github.com/labstack/echo/v4"
)

type NewsController struct {
	NewsRepo news.Repository
}

func NewNewsController(newsRepo news.Repository) *NewsController {
	return &NewsController{
		NewsRepo: newsRepo,
	}
}

func (newsController NewsController) GetNewsByCategory(c echo.Context) error {
	category := "otomotif"
	fmt.Println(category)
	data, error := newsController.NewsRepo.GetNewsByCategory(category)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccessResponse(c, FromDomain(data))
}
