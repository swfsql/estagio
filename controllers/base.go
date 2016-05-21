package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	//this.Layout = "layout.html"
	//this.LayoutSections = make(map[string]string)
}

func (this *BaseController) All() {
	var all map[interface{}]interface{}
	all = make(map[interface{}]interface{})
	this.Data["all"] = all

	for key, value := range this.Data {
		if key != "all" {
			fmt.Println(key)
			fmt.Println(value)
			all[key] = value
		}
	}
}
