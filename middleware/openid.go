package middleware

import (
	"errors"
	"github.com/labstack/echo"
	"golang-echo-layout/model"
	"golang-echo-layout/service"
	"golang-echo-layout/utils"
	"net/http"
)

type UserContext struct {
	echo.Context
}

func (u *UserContext) UserInfo() (*model.User, error) {
	userService := service.NewUserService()
	user := userService.GetUserByOpenId(u.Request().Header.Get("open-id"))
	if user.ID == 0 {
		return nil, errors.New("Unauthorized")
	}

	return user, nil
}

func ValidateOpenId() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if len(c.Request().Header.Get("open-id")) == 0 {
				return utils.NewHTTPError(http.StatusUnauthorized, "账号验证不通过，请重新允许小程序授权", nil)
			}
			cc := &UserContext{c}
			user, err := cc.UserInfo()
			if err != nil {
				return utils.NewHTTPError(http.StatusUnauthorized, "账号验证不通过，请重新允许小程序授权!", nil)
			}

			cc.Context.Set("user_id", user.ID)
			return next(cc)
		}
	}
}
