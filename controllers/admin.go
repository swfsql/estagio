package controllers

import (
	"fmt"
	"github.com/swfsql/estagio/models"
)

type AdminController struct {
	AuthController
}

func (this *AdminController) Get() {
	sess := this.StartSession()
	conta := sess.Get("conta").(models.Conta)

	fmt.Println("ENTROU NA PAGINA DO Admin")

	switch conta.Pessoa.Privilegio {
	case 5:
		this.TplName = "admin.html"
		//estagios, _ := sess.Get("estagios").([]*models.Estagio)
		//this.Data["Estagios"] = estagios
		break
	default:
		this.Redirect("/", 302)
		return
	}

	this.Data["Usuario"] = conta.Pessoa
	fmt.Println(conta.Pessoa.Privilegio)
	fmt.Println(conta.Pessoa.Nome)

	fmt.Println(this.TplName)
	this.All()
	this.Render()
}
