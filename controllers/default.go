package controllers

import (
	"fmt"
	"github.com/swfsql/estagio/models"
	"strconv"
)

type IndexController struct {
	AuthController
}

func (this *IndexController) Get() {
	sess := this.StartSession()
	conta := sess.Get("conta").(models.Conta)

	fmt.Println("Entrou no DEFAULT")
	this.Data["Usuario"] = conta.Pessoa
	fmt.Println(conta.Pessoa.Privilegio)
	fmt.Println(conta.Pessoa.Nome)

	switch conta.Pessoa.Privilegio {
	case 2: // aluno
		aluno, _ := sess.Get("aluno").(models.Aluno)
		this.Redirect("/discente/"+strconv.FormatUint(aluno.Id, 10), 302)
		return
		break
	case 3: // professor normal
		professor, _ := sess.Get("professor").(models.Professor)
		this.Redirect("/docente/"+strconv.FormatUint(professor.Id, 10), 302)
		/*estagios, _ := sess.Get("estagios").([]*models.Estagio)
		this.TplName = "professor.html"
		this.Data["Professor"] = professor
		this.Data["Curso"] = professor.Curso
		this.Data["Estagios"] = estagios*/
		return
		break
	case 4: //professor coordenador
		professor, _ := sess.Get("professor").(models.Professor)
		this.Redirect("/docente/"+strconv.FormatUint(professor.Id, 10), 302)
		/*estagios, _ := sess.Get("estagios").([]*models.Estagio)
		this.TplName = "professor.html"
		this.Data["Professor"] = professor
		this.Data["Curso"] = professor.Curso
		this.Data["Estagios"] = estagios*/
		return
		break
	case 5: // admin
		this.Redirect("/nucleo/", 302)
		return
	}

	fmt.Println(this.TplName)
	this.All()
	this.Render()
}
