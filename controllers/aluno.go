package controllers

import (
	"fmt"
	"github.com/swfsql/estagio/models"
	"strconv"
)

type AlunoController struct {
	AuthController
}

func (this *AlunoController) Get() {
	sess := this.StartSession()
	conta := sess.Get("conta").(models.Conta)

	aluno_id, _ := strconv.ParseUint(this.Ctx.Input.Param(":id"), 10, 64)
	fmt.Println("ENTROU NA PAGINA DO ALUNO")

	switch conta.Pessoa.Privilegio {
	case 2: // aluno
		aluno, _ := sess.Get("aluno").(models.Aluno)
		if aluno_id != aluno.Id {

			this.Redirect("/", 302)
		}
		estagios, _ := sess.Get("estagios").([]*models.Estagio)
		this.TplName = "aluno.html"
		this.Data["Aluno"] = aluno
		this.Data["Curso"] = aluno.Curso
		this.Data["Estagios"] = estagios
		break
	case 3: // professor normal
		break
	case 4: //professor coordenador
		break
	case 5: // admin
		this.TplName = "aluno_admin.html"
		break
	}

	this.Data["Usuario"] = conta.Pessoa
	fmt.Println(conta.Pessoa.Privilegio)
	fmt.Println(conta.Pessoa.Nome)

	fmt.Println(this.TplName)
	this.All()
	this.Render()
}