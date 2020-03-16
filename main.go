package main

import (
	"demo-backend/command"
	_ "demo-backend/routers"

	"github.com/astaxie/beego"
	_ "github.com/keemis/library/flag"
)

func main() {

	command.Run()

	beego.Run()
}
