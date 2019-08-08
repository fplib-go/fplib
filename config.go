package fplib

import (
	"github.com/astaxie/beego"
)

var (
	is_Debug bool
)

func init() {

	if beego.AppConfig.String("runmode") == "prod" {
		is_Debug = false
	} else {
		is_Debug = true
	}

}
