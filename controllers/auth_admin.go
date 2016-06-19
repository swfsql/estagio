package controllers

import (
	"fmt"
	"github.com/swfsql/estagio/models"
)

type AuthAdminController struct {
	BaseController
}

func (this *AuthAdminController) Prepare() {
	sess := this.StartSession()
	conta, achou := sess.Get("conta").(models.Conta)

	fmt.Println("entrou no AUTH")

	if achou == false || conta.Pessoa.Privilegio != 5 { // nao esta logado
		fmt.Println("mas nao passou no AUTH")
		this.Redirect("/login", 302)
	}
}
