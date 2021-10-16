package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/hanifbg/login_register/config"
	user "github.com/hanifbg/login_register/entity/user"

	"github.com/hanifbg/login_register/handler"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.HTTPErrorHandler = handler.ErrorHandler
	var dbUser, dbPass, dbHost, dbPort, dbName, port string

	//env setting

	cfg := config.ProviderConfig()
	port = cfg.GetString("server.port")
	dbUser = cfg.GetString("database.username")
	dbPass = cfg.GetString("database.password")
	dbHost = cfg.GetString("database.host")
	dbPort = cfg.GetString("database.port")
	dbName = cfg.GetString("database.dbname")

	//DB connection
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}

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
