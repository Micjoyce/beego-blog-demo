package controllers

import (
	"fmt"
	"myblog/models"

	"github.com/astaxie/beego"
)

// CategoryController xx
type CategoryController struct {
	beego.Controller
}

// Get xx
func (c *CategoryController) Get() {
	op := c.Input().Get("op")
	switch op {
	case "add":
		name := c.Input().Get("category")
		if len(name) == 0 {
			break
		}
		// 写数据
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 301)
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		// 删除数据
		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 301)
		return
	}
	// c.Catetories =
	c.Data["IsCategory"] = true
	c.TplName = "category.html"

	var err error
	c.Data["Categories"], err = models.GetAllCategories()
	fmt.Println(c.Data["Categories"])
	if err != nil {
		beego.Error(err)
	}
}
