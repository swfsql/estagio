package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/swfsql/estagio/models"
)

type MyController struct {
	beego.Controller
}

func (c *MyController) Get() {
	c.TplName = "a/index.tpl"

	o := orm.NewOrm()
	o.Insert(&models.Curso{Nome: "curso01"})
	o.Insert(&models.Curso{Nome: "curso02"})
	o.Insert(&models.Curso{Nome: "curso03"})

	o.Insert(&models.Pessoa{Nome: "aluno01"})
	o.Insert(&models.Pessoa{Nome: "aluno02"})
	o.Insert(&models.Pessoa{Nome: "aluno03"})

	o.Insert(&models.Pessoa{Nome: "prof01"})
	o.Insert(&models.Pessoa{Nome: "prof02"})
	o.Insert(&models.Pessoa{Nome: "prof03"})

	//pessoa := models.Pessoa{Id: 0, Nome: "user01"}

	var pessoas []*models.Pessoa
	qs := o.QueryTable("pessoa")
	qs.All(&pessoas)

	var cursos []*models.Curso
	qs = o.QueryTable("curso")
	qs.All(&cursos)

	o.Insert(&models.Aluno{Pessoa: pessoas[0], Curso: cursos[0]})
	o.Insert(&models.Aluno{Pessoa: pessoas[1], Curso: cursos[0]})
	o.Insert(&models.Aluno{Pessoa: pessoas[2], Curso: cursos[1]})

	o.Insert(&models.Prof{Pessoa: pessoas[3], Curso: cursos[0]})
	o.Insert(&models.Prof{Pessoa: pessoas[4], Curso: cursos[0]})
	o.Insert(&models.Prof{Pessoa: pessoas[5], Curso: cursos[1]})

	var alunos []*models.Aluno
	qs = o.QueryTable("aluno")
	qs.All(&alunos)

	var profs []*models.Prof
	qs = o.QueryTable("prof")
	qs.All(&profs)

	o.Insert(&models.Estagio{Aluno: alunos[0], Prof: profs[0]})

	var estagios []*models.Estagio
	qs = o.QueryTable("estagio")
	qs.All(&estagios)

	for _, p := range pessoas {
		fmt.Println(p)
	}
	for _, p := range alunos {
		fmt.Println(p)
	}
	for _, p := range profs {
		fmt.Println(p)
	}
	for _, p := range estagios {
		fmt.Println(p)
	}

	//p := models.Pessoa{Id: pessoa.Id}
	//o.Read(&p)
	//fmt.Printf("pessoa lida: ")
	//fmt.Println(p)

}

func (c *MyController) Post() {
	c.TplName = "a/b.tpl"
}
