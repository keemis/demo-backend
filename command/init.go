package command

import (
	"github.com/keemis/library/logs"
	"github.com/keemis/library/mysql"
)

func init() {

	logs.Init(nil)

	mysql.Init()

}
