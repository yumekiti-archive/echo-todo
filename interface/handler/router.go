package handler

import (
	"github.com/labstack/echo/v4"
)

// InitRouting routesの初期化
func InitRouting(
	e *echo.Echo,
	taskHandler TaskHandler,
) {
	e.GET("/task", taskHandler.GetAll())
	e.GET("/task/:id", taskHandler.Get())
	e.POST("/task", taskHandler.Post())
	e.PUT("/task/:id", taskHandler.Put())
	e.DELETE("/task/:id", taskHandler.Delete())
}
