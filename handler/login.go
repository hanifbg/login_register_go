package handler

import (
	"net/http"

	user "github.com/hanifbg/login_register/entity/user"
	echo "github.com/labstack/echo/v4"
)

func LoginHandler(c echo.Context) (err error) {
	cc := c.(Option)
	l := new(user.LoginUser)
	if err = c.Bind(l); err != nil {
		return
	}

	result, err := cc.Srv.UserService.LoginUser(*l)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
