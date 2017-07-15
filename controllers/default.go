package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

type MainController struct {
	BaseController
}

func init() {
	beego.SetStaticPath("/assets", "assets")
	//beego.BConfig.WebConfig.DirectoryIndex = true
}
func (this *BaseController) SetLoginInfo(membercode string) {
	loginAccount := this.GetSession("loginAccount")
	if loginAccount == nil {
		this.SetSession("loginAccount", membercode)
	}
}

func (this *BaseController) VerifyLogin() bool {
	loginAccount := this.GetSession("loginAccount")
	if loginAccount == nil {
		this.Redirect("/Login", 302)
		return false
	} else {
		return true
	}
}

func (c *MainController) Get() {
	c.Data["Website"] = "billwang"
	c.Data["Email"] = "billwang103@gmail.com"
	c.TplName = "login.html" //index.tpl
}
