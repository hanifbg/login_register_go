package handler

import (
	"github.com/hanifbg/login_register/service"
	echo "github.com/labstack/echo/v4"
)

type HandlerContext struct {
	echo.Context
	Srv *service.Services
}
