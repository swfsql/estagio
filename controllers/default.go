package controllers

import (
	"fmt"
	"github.com/swfsql/estagio/models"
)

type IndexController struct {
	AuthController
}

func (this *IndexController) Get() {
	sess := this.StartSession()
	conta := sess.Get("conta").(models.Conta)

	this.Data["Usuario"] = conta.Pessoa
	fmt.Println(conta.Pessoa.Privilegio)
	fmt.Println(conta.Pessoa.Nome)

	switch conta.Pessoa.Privilegio {
	case 2:
		aluno, _ := sess.Get("aluno").(models.Aluno)
		estagios, _ := sess.Get("estagios").([]*models.Estagio)

		this.TplName = "aluno.html"
		this.Data["HeadTitle"] = "Aluno Title"
		this.Data["HeadScripts"] = []string{
			"/static/js/jquery.min.js",
			"/static/js/aluno.js"}
		this.Data["HeadStyles"] = []string{
			"/static/css/bootstrap-theme.min.css",
			"/static/css/bootstrap.min.css",
			"/static/css/aluno.css"}

		this.Data["Aluno"] = aluno
		this.Data["Curso"] = aluno.Curso
		this.Data["Estagios"] = estagios
		break
	case 3:
		this.TplName = "index_professor.html"
		break
	case 4:
		this.TplName = "index_coord.html"
		break
	case 5:
		this.TplName = "index_admin.html"
		break
	}

	fmt.Println(this.TplName)
	this.All()
	this.Render()
}
