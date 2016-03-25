package main

import (
	"github.com/astaxie/beego"
	_ "github.com/swfsql/estagio/routers"
)

func main() {
	beego.Run()
}
