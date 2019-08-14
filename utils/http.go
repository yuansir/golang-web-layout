package utils

import (
	"github.com/labstack/echo"
	"golang-echo-layout/utils/log"
	"net/http"
)

type (
	HTTPError struct {
		Code    int      `json:"code"`
		Message string   `json:"message"`
		Errors  []string `json:"errors"`
	}
	HTTPSuccess struct {
		Code    int                    `json:"code"`
		Message string                 `json:"message"`
		Data    map[string]interface{} `json:"data"`
	}
)

func (e *HTTPError) Error() string {
	return e.Message
}

func NewHTTPError(code int, message string, errors []string) *HTTPError {
	return &HTTPError{Code: code, Message: message, Errors: errors}
}

func NewHTTPSuccess(c echo.Context, message string, data map[string]interface{}) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": message,
		"data":    data,
	})
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		msg  string
		errs []string
	)

	if he, ok := err.(*HTTPError); ok {
		code = he.Code
		msg = he.Message
		errs = he.Errors
	} else if ee, ok := err.(*echo.HTTPError); ok {
		code = ee.Code
		msg = ee.Message.(string)
		errs = SplitErrors(err)
	} else {
		msg = err.Error()
	}

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, map[string]interface{}{
				"code":    code,
				"message": msg,
				"errors":  errs,
			})
		}
		if err != nil {
			log.Error(err)
		}
	}
}
