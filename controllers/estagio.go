package controllers

import (
	"fmt"
	"strconv"

	"github.com/swfsql/estagio/models"
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

		estagioIndice := -1
		for i, e := range estagios {
			if estagio_id == e.Id {
				estagioIndice = i
				break
			}
		}
		if estagioIndice == -1 {
			this.Redirect("/", 302)
			return
		}
		estagio := estagios[estagioIndice]

		this.TplName = "estagio_aluno.html"
		this.Data["Aluno"] = aluno
		this.Data["Curso"] = aluno.Curso
		this.Data["Estagio"] = estagio
		break
	case 3, 4: // professor normal
		//professor, _ := sess.Get("professor").(models.Professor)
		estagios, _ := sess.Get("estagios").([]*models.Estagio)

		estagioIndice := -1

		for i, e := range estagios {
			if estagio_id == e.Id {
				estagioIndice = i
				break
			}
		}

		if estagioIndice == -1 {
			this.Redirect("/", 302)
			return
		}

		estagio := estagios[estagioIndice]

		this.TplName = "estagio_professor.html"
		//	this.Data["Professor"] = professor
		//	this.Data["Curso"] = professor.Curso
		this.Data["Estagio"] = estagio
		break

		/*	case 4: //professor coordenador
			break */

	case 5: // admin
		estagios, _ := models.GetEstagios()

		estagioIndice := -1

		for i, e := range estagios {
			if estagio_id == e.Id {
				estagioIndice = i
				break
			}
		}

		if estagioIndice == -1 {
			this.Redirect("/", 302)
			return
		}

		fmt.Println(estagioIndice)
		estagio := estagios[estagioIndice]

		this.TplName = "estagio_admin.html"
		this.Data["Estagio"] = estagio
		//this.Data[]
		break
	}

	this.Data["Usuario"] = conta.Pessoa
	fmt.Println(conta.Pessoa.Privilegio)
	fmt.Println(conta.Pessoa.Nome)

	fmt.Println(this.TplName)
	this.All()
	this.Render()
}
