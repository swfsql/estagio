package controllers

import (
	"github.com/swfsql/estagio/models"
)

type AuthController struct {
	BaseController
}

func (this *AuthController) Prepare() {
	this.BaseController.Prepare()
	sess := this.StartSession()
	//defer sess.SessionRelease()
	if _, achou := sess.Get("conta").(models.Conta); achou == false { // nao esta logado
		this.Redirect("/login", 302)
	}
}
