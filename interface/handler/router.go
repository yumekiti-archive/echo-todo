package handler

import (
	"github.com/labstack/echo/v4"
)

// InitRouting routesの初期化
func InitRouting(
	e *echo.Echo,
	taskHandler TaskHandler,
) {
	e.GET("/task", taskHandler.Get())
}
