package controllers

import "goutils"

type BetsuggestController struct {
	BaseController
}

func (this *BetsuggestController) Get() {
	this.VerifyLogin()
	BUCode := this.Input().Get("BUCode")
	this.Data["BUCode"] = BUCode
	goutils.Logger.Info("BUCode:" + BUCode)
	this.TplName = "betsuggest.html"
}
