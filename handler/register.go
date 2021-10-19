package handler

import (
	"net/http"

	user "github.com/hanifbg/login_register/entity/user"
	echo "github.com/labstack/echo/v4"
)

func RegisterHandler(c echo.Context) error {
	cc := c.(Option)
	u := new(user.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := c.Validate(u); err != nil {
		return err
	}

	user := cc.Srv.UserService.BindUser(*u)

	return c.JSON(http.StatusOK, user)
}
