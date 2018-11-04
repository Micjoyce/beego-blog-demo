package controllers

import (
	"github.com/astaxie/beego"
)

// MainController xxxx
type MainController struct {
	beego.Controller
}

// Get home get
func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "index.html"
}
