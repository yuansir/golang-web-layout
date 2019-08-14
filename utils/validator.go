package utils

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	"net/http"
)

//func init() {
//
//}

func ValidateRequestStrut(c echo.Context, data interface{}) *HTTPError {
	if err := c.Bind(data); err != nil {
		return NewHTTPError(http.StatusBadRequest, "参数验证失败", SplitErrors(err))
	}

	if _, err := govalidator.ValidateStruct(data); err != nil {
		return NewHTTPError(http.StatusBadRequest, "参数验证失败", SplitErrors(err))
	}

	return nil
}
