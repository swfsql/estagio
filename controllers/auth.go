package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/swfsql/estagio/models"
)

type AuthController struct {
	beego.Controller
}

func (this *AuthController) Prepare() {
	sess := this.StartSession()
	//defer sess.SessionRelease()

	fmt.Println("~~~~~~~~~~~~~~~~~ ENTROU NO PREPARE ~~~~~~~~~~~~~~~~~~")

	this.Layout = "layout.html"
	this.LayoutSections = make(map[string]string)

	if _, achou := sess.Get("conta").(models.Conta); achou == false { // nao esta logado
		this.Redirect("/login", 302)
	}
	fmt.Println("~~~~~~~~~~~~~~~~~ SAIU DO PREPARE ~~~~~~~~~~~~~~~~~~")
}

func (this *AuthController) All() {
	fmt.Println("~~~~~~~~~~~~~~~~~ ENTROU NO FINISH ~~~~~~~~~~~~~~~~~~")
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
	fmt.Println("~~~~~~~~~~~~~~~~~ SAIU DO FINISH ~~~~~~~~~~~~~~~~~~")
}
