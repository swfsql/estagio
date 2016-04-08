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

	this.Data["usuario"] = conta.Pessoa
	fmt.Println(conta.Pessoa.Privilegio)
	fmt.Println(conta.Pessoa.Nome)

	switch conta.Pessoa.Privilegio {
	case 2:
		this.TplName = "index_aluno.html"
		aluno, _ := sess.Get("aluno").(models.Aluno)
		estagios, _ := sess.Get("estagios").([]*models.Estagio)
		this.TplName = "index_aluno.html"
		fmt.Println("~~~~~~~ok 01~~~~~~~")
		this.Data["aluno"] = aluno
		this.Data["curso"] = aluno.Curso
		this.Data["estagios"] = estagios
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
