package main

import (
	"demo-backend/command"
	_ "demo-backend/routers" // 注入路由

	"github.com/astaxie/beego"
	_ "github.com/keemis/library/flag" // 注入版本号
)

func main() {

	command.Run()

	beego.Run()
}
