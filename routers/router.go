package routers

import (
	"github.com/astaxie/beego"
	"github.com/swfsql/estagio/controllers"
)

func init() {
	beego.Router("/", &controllers.MyController{})

}
