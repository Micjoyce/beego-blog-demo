package main

import (
	"myblog/models"
	_ "myblog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	o := orm.NewOrm()
	o.Using("default")
	beego.Run()
}
