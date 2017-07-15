package controllers

import (
	"log"

	"qqpaygc/models"

	_ "github.com/go-sql-driver/mysql"
)

// Operations about object
//改了Routing 和controler的名稱 要跑過Bee Run在launch才會生效
type LoginController struct {
	BaseController
}

//
// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (this *LoginController) Post() {
	var inputObj models.APILoginInput
	log.Println("LoginController RequestBody:", string(this.Ctx.Input.RequestBody))

	inputObj.Membercode = this.Input().Get("membercode")
	inputObj.Password = this.Input().Get("password")
	//json.Unmarshal(this.Ctx.Input.RequestBody, &inputObj) //把JSON值塞進Object去
	//Login(驗證)

	result := true //storage.DB_BOLoginVerify(inputObj.Membercode, inputObj.Password)
	log.Println("verifyLogin result>>:", result)
	if result {
		//登入成功
		//this.TplName = "index.html"
		this.SetLoginInfo(inputObj.Membercode)
		this.Redirect("/Index", 302)
	} else {
		//登入失敗
		//this.TplName = "login.html"
		this.Redirect("/Login", 302)
	}

	this.TplName = "login.html"

}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}
