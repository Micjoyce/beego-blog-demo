package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	beecontext "github.com/astaxie/beego/context"
)

// LoginController login controller
type LoginController struct {
	beego.Controller
}

// Get xxx
func (c *LoginController) Get() {
	fmt.Println(c.Input().Get("exit"))
	isExit := c.Input().Get("exit") == "true"
	if isExit {
		c.Ctx.SetCookie("username", "", 0, "/")
		c.Ctx.SetCookie("password", "", 0, "/")
		c.Redirect("/", 302)
		return
	}
	c.TplName = "login.html"
}

// Post xxx
func (c *LoginController) Post() {
	username := c.Input().Get("username")
	password := c.Input().Get("password")
	autoLogin := c.Input().Get("autoLogin")
	if beego.AppConfig.String("username") == username &&
		beego.AppConfig.String("password") == password {
		maxAge := 0
		if autoLogin == "on" {
			maxAge = 1<<32 - 1
		}
		c.Ctx.SetCookie("username", username, maxAge, "/")
		c.Ctx.SetCookie("password", password, maxAge, "/")
	}
	c.Redirect("/", 301)
	return
}

func checkAccount(ctx *beecontext.Context) bool {
	ck, err := ctx.Request.Cookie("username")
	if err != nil {
		return false
	}
	username := ck.Value
	if username != beego.AppConfig.String("username") {
		return false
	}
	ck, err = ctx.Request.Cookie("password")
	if err != nil {
		return false
	}
	password := ck.Value
	if password != beego.AppConfig.String("password") {
		return false
	}
	fmt.Println("--------------------")
	return true
}
