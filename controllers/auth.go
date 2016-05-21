package controllers

import (
	"fmt"
	"github.com/swfsql/estagio/models"
)

type AuthController struct {
	BaseController
}

func (this *AuthController) Prepare() {
	//this.BaseController.Prepare()
	sess := this.StartSession()
	//defer sess.SessionRelease()

	fmt.Println("entrou no AUTH")

	if _, achou := sess.Get("conta").(models.Conta); achou == false { // nao esta logado
		fmt.Println("mas nao passou no AUTH")
		this.Redirect("/login", 302)
	}
}
