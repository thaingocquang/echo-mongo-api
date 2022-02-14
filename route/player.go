package route

import (
	"TestGoAPI/controller"
	"TestGoAPI/validation"
	"github.com/labstack/echo/v4"
)

func Player(e *echo.Echo) {
	players := e.Group("/players")
	players.POST("", controller.PlayerCreate, validation.PlayerCreate)
	players.GET("/:id", controller.PlayerFindByID, validation.PlayerValidateID, validation.PlayerCheckExistedByID)
}
