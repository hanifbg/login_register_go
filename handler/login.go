package handler

import (
	"net/http"

	user "github.com/hanifbg/login_register/entity/user"
	echo "github.com/labstack/echo/v4"
)

func LoginHandler(c echo.Context) (err error) {
	cc := c.(HandlerContext)
	l := new(user.LoginUser)
	if err = c.Bind(l); err != nil {
		return
	}

	result := cc.Srv.UserService.LoginUser(*l)
	return c.JSON(http.StatusOK, result)
}
