package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/swfsql/estagio/models"
)

type ListaPessoaController struct {
	beego.Controller
}

func (this *ListaPessoaController) PegarPessoas() {

	o := orm.NewOrm()

	var pessoas []*models.Pessoa
	qs := o.QueryTable("pessoa")
	qs.All(&pessoas)

	dado := struct{ Pessoas []*models.Pessoa }{pessoas}
	this.Data["json"] = dado
	this.ServeJSON()

}

func (this *ListaPessoaController) CadastrarPessoa() {

	o := orm.NewOrm()
	fmt.Println("hauauauauaua")

	dado := struct {
		Nome     string
		Telefone string
		Email    string
	}{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &dado)
	if err != nil {
		fmt.Println(err)

		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("JSON invalido"))

		return
	}
	o.Insert(&models.Pessoa{Nome: dado.Nome, Email: dado.Email, Telefone: dado.Telefone, Senha: dado.Nome + "0", Cadastro: dado.Email})
	//fmt.Println(dado.Nome)
}
