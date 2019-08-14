package handler

import (
	"github.com/labstack/echo"
)

type BaseHandler struct {
}

func (b *BaseHandler) GetUserId(c echo.Context) uint {
	return c.Get("user_id").(uint)
}
