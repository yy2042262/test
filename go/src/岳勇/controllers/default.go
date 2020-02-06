package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	/*c.Data["Website"] = "岳勇"
	c.Data["Email"] = "yueyong@cqsf.com"
	*/
	name:=c.GetString("name")
	email:=c.GetString("email")
	c.Data["Website"]=name
	c.Data["Email"]=email
	c.TplName = "index.tpl"
}
