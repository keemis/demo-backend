package routers

import (
	"demo-backend/controller"

	"github.com/astaxie/beego"
)

func init() {

	ns := beego.NewNamespace("/api",

		// 学校接口
		beego.NSNamespace("/school",
			beego.NSRouter("/list", &controller.SchoolController{}, "GET:GetSchools"), // 学校列表
		),

	)

	beego.AddNamespace(ns)
}
