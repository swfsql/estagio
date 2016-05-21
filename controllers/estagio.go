package controllers

import (
	"fmt"
	"github.com/swfsql/estagio/models"
	"strconv"
)

type EstagioController struct {
	AuthController
}

func (this *EstagioController) Get() {
	sess := this.StartSession()
	conta := sess.Get("conta").(models.Conta)

	estagio_id, _ := strconv.ParseUint(this.Ctx.Input.Param(":id"), 10, 64)
	fmt.Println("ENTROU NA PAGINA DO ESTAGIO")

	switch conta.Pessoa.Privilegio {
	case 2: // aluno
		aluno, _ := sess.Get("aluno").(models.Aluno)
		estagios, _ := sess.Get("estagios").([]*models.Estagio)

		estagio_indice := -1
		for i, e := range estagios {
			if estagio_id == e.Id {
				estagio_indice = i
				break
			}
		}
		if estagio_indice == -1 {
			this.Redirect("/", 302)
			return
		}
		estagio := estagios[estagio_indice]

		this.TplName = "estagio_aluno.html"
		this.Data["Aluno"] = aluno
		this.Data["Curso"] = aluno.Curso
		this.Data["Estagio"] = estagio
		break
	case 3: // professor normal
		break
	case 4: //professor coordenador
		break
	case 5: // admin
		this.TplName = "estagio_admin.html"
		break
	}

	this.Data["Usuario"] = conta.Pessoa
	fmt.Println(conta.Pessoa.Privilegio)
	fmt.Println(conta.Pessoa.Nome)

	fmt.Println(this.TplName)
	this.All()
	this.Render()
}
