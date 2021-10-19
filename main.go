package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/hanifbg/login_register/config"
	driver "github.com/hanifbg/login_register/driver"
	user "github.com/hanifbg/login_register/entity/user"
	"github.com/hanifbg/login_register/repository"
	service "github.com/hanifbg/login_register/service"

	"github.com/hanifbg/login_register/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.HTTPErrorHandler = handler.ErrorHandler
	e.Debug = true

	//env setting
	cfg := config.ProviderConfig()
	port := cfg.GetString("server.port")

	//DB
	db := driver.DBConnection(cfg)
	db.AutoMigrate(&user.User{})

	repo := initRepository(repository.Option{
		DB: db,
	})
	srv := initService(service.Option{
		Repository: repo,
	})

	//routes
	v1 := e.Group("/v1", func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := handler.Option{c, srv}
			return h(cc)
		}
	})
	v1.POST("/register", handler.RegisterHandler)
	v1.POST("/login", handler.LoginHandler)

	e.Logger.Fatal(e.Start(":" + port))
}

func initRepository(repoOption repository.Option) *repository.Repository {
	userRepo := repository.NewUserRepository(repoOption)

	repo := repository.Repository{
		User: userRepo,
	}

	return &repo
}

func initService(servOption service.Option) *service.Services {
	us := service.NewUserService(servOption)

	srv := service.Services{
		UserService: us,
	}

	return &srv
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
