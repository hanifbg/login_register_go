package handler

import (
	"net/http"

	user "github.com/hanifbg/login_register/entity/user"
	echo "github.com/labstack/echo/v4"
)

func RegisterHandler(c echo.Context) (err error) {
	u := new(user.User)
	if err = c.Bind(u); err != nil {
		return
	}
	if err = c.Validate(u); err != nil {
		return
	}

	return c.JSON(http.StatusOK, u)
}
