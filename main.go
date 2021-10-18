package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/hanifbg/login_register/config"
	user "github.com/hanifbg/login_register/entity/user"
	"github.com/hanifbg/login_register/repository"
	service "github.com/hanifbg/login_register/service"
	"gorm.io/gorm"

	"github.com/hanifbg/login_register/handler"
	"github.com/labstack/echo/v4"
)

type SrvContext struct {
	echo.Context
	srv *service.Services
}

func RegisterHandler(c echo.Context) error {
	cc := c.(*SrvContext)
	u := new(user.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := c.Validate(u); err != nil {
		return err
	}

	user := cc.srv.UserService.BindUser(*u)

	return c.JSON(http.StatusOK, user)
}

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
	repo := initRepository(db)
	srv := initService(service.Option{
		Repository: repo,
	})
	v1 := e.Group("/v1", func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := handler.HandlerContext{c, srv}
			return h(cc)
		}
	})
	v1.POST("/register", handler.RegisterHandler)
	v1.POST("/login", handler.LoginHandler)

	e.Logger.Fatal(e.Start(":" + port))
}

func initRepository(db *gorm.DB) *repository.Repository {
	userRepo := repository.NewUserRepository(db)

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
