package controllers

import (
	"fmt"
	"strconv"

	"github.com/swfsql/estagio/models"
)

type ProfessorController struct {
	AuthController
}

func (this *ProfessorController) Get() {
	sess := this.StartSession()
	conta := sess.Get("conta").(models.Conta)

	professor_id, _ := strconv.ParseUint(this.Ctx.Input.Param(":id"), 10, 64)
	fmt.Println("ENTROU NA PAGINA DO PROFESSOR")

	switch conta.Pessoa.Privilegio {
	case 2: // aluno
		this.Redirect("/", 302)
		return
	case 3: // professor normal
		professor, _ := sess.Get("professor").(models.Professor)
		if professor_id != professor.Id {
			this.Redirect("/", 302)
			return
		}

		estagios, _ := sess.Get("estagios").([]*models.Estagio)
		this.TplName = "professor.html"
		this.Data["Professor"] = professor
		this.Data["Curso"] = professor.Curso
		this.Data["Estagios"] = estagios
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
