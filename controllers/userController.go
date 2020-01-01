package controllers

import (
	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

type userController struct {
	mvc.Controller
}

func init() {
	mvc.AddHandlers(&userController{})
}

func (user *userController) Login() {
	user.Response.RemoveSession()
	user.UseTplPath()
}

type userInfo struct {
	Code     string `json:"code"`
	Pwd      string `json:"pwd"`
	UserName string `json:"username"`
}

func (user *userController) DoLoginPost(u *userInfo) {
	tmp := struct {
		UserName string `json:"username"`
		IsLogin  bool   `json:"islogin"`
	}{IsLogin: false}
	if u.UserName == "huzhouyu" && u.Pwd == "123456" {
		user.Response.SetSession("userinfo", u)
		tmp.UserName = u.UserName
		tmp.IsLogin = true
	}
	user.Response.WriteJson(tmp)
}
func (user *userController) VueTest() {
	user.UseTplPath()
}
