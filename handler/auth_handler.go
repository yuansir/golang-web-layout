package handler

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req"
	"github.com/labstack/echo"
	"golang-echo-layout/config"
	"golang-echo-layout/repository"
	"golang-echo-layout/utils"
	"golang-echo-layout/utils/log"
	"net/http"
)

type AuthHandler struct {
}

type Code2SessionResp struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

func (a *AuthHandler) JSCode2Session(c echo.Context) error {
	r, err := req.Get(fmt.Sprintf("%s&appid=%s&secret=%s&js_code=%s",
		utils.Code2SessionUrl,
		config.Conf.Wx.AppId,
		config.Conf.Wx.AppSecret,
		c.QueryParam("code"), ))

	if err != nil {
		log.Error(err)
		return utils.NewHTTPError(http.StatusBadRequest, err.Error(), nil)
	}

	if r.Response().StatusCode != http.StatusOK {
		return utils.NewHTTPError(r.Response().StatusCode, "", nil)
	}

	var response Code2SessionResp
	respStr, _ := r.ToString()
	_ = json.Unmarshal([]byte(respStr), &response)
	if response.Errcode != 0 {
		return utils.NewHTTPError(http.StatusBadRequest, response.Errmsg, nil)
	}

	userRepo := repository.NewUserRepository()
	userRepo.FirstOrCreate(response.OpenId, response.Unionid, response.SessionKey)
	return utils.NewHTTPSuccess(c, "OK", map[string]interface{}{
		"openid": response.OpenId,
	})
}
