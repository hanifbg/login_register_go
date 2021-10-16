package handler

import (
	"net/http"
	"time"

	user "github.com/hanifbg/login_register/entity/user"
	echo "github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(c echo.Context) (err error) {
	u := new(user.User)
	if err = c.Bind(u); err != nil {
		return
	}
	if err = c.Validate(u); err != nil {
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 8)

	user := user.User{}
	user.Name = u.Name
	user.Email = u.Email
	user.Phone_number = u.Phone_number
	user.Address = u.Address
	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, hashedPassword)
}
