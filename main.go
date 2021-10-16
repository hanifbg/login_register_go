package main

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	user "github.com/hanifbg/login_register/entity/user"

	"github.com/hanifbg/login_register/handler"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.HTTPErrorHandler = handler.ErrorHandler
	var dbUser, dbPass, dbHost, dbPort, dbName, port string

	//env setting
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.config")
	err := viper.ReadInConfig()
	if err != nil {
		port = os.Getenv("PORT")
		dbUser = os.Getenv("dbUsername")
		dbPass = os.Getenv("dbPassword")
		dbHost = os.Getenv("dbHost")
		dbPort = "3306"
		dbName = os.Getenv("dbName")
	} else {
		port = viper.GetString("server.port")
		if port == "" {
			port = os.Getenv("PORT")
		}

		//config db
		if viper.GetString("database.username") == "" {
			dbUser = os.Getenv("dbUsername")
			dbPass = os.Getenv("dbPassword")
			dbHost = os.Getenv("dbHost")
			dbPort = "3306"
			dbName = os.Getenv("dbName")
		} else {
			dbUser = viper.GetString("database.username")
			dbPass = viper.GetString("database.password")
			dbHost = viper.GetString("database.host")
			dbPort = viper.GetString("database.port")
			dbName = viper.GetString("database.dbname")
		}
	}

	//DB connection
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}

	db.AutoMigrate(&user.User{})

	//routes
	e.POST("/login", handler.LoginHandler)
	e.POST("/register", handler.RegisterHandler)

	e.Logger.Fatal(e.Start(":" + port))
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
