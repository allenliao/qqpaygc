package controllers

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	this.VerifyLogin()
	this.TplName = "index.html" //index.tpl
}
