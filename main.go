package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/hanifbg/login_register/config"
	user "github.com/hanifbg/login_register/entity/user"
	"github.com/hanifbg/login_register/repository"

	"github.com/hanifbg/login_register/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.HTTPErrorHandler = handler.ErrorHandler

	//env setting

	cfg := config.ProviderConfig()
	port := cfg.GetString("server.port")
	db := repository.Connection(cfg)

	db.AutoMigrate(&user.User{})

	//routes
	v1 := e.Group("/v1")
	groupV1Routes(v1)

	e.Logger.Fatal(e.Start(":" + port))
}

func groupV1Routes(e *echo.Group) {
	e.POST("/login", handler.LoginHandler)
	e.POST("/register", handler.RegisterHandler)
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
