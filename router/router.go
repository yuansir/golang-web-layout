package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang-echo-layout/handler"
	myMiddleware "golang-echo-layout/middleware"
	"golang-echo-layout/service"
	"golang-echo-layout/utils"
)

func Routers() *echo.Echo {
	e := echo.New()
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "static",
		Browse: true,
	}))

	e.HTTPErrorHandler = utils.CustomHTTPErrorHandler
	e.Use(myMiddleware.Logger())
	e.Use(middleware.Recover())
	api := e.Group("/api")

	authHandler := handler.AuthHandler{}
	api.GET("/code2session", authHandler.JSCode2Session)

	return e
}