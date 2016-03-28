package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/swfsql/estagio/models"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	sess := this.StartSession()
	//defer sess.SessionRelease()

	if conta, achou := sess.Get("conta").(models.Conta); achou == false { // nao esta logado
		//this.TplName = "login.html"
		//this.Ctx.Redirect(302, "/login")
		this.Redirect("/login", 302)
		return
	} else { // esta logado
		//pessoa := models.Pessoa{Nome: pessoa.Nome, Telefone: pessoa.Telefone, Email: pessoa.Email, Privilegio: pessoa.Privilegio}
		this.Data["pessoa"] = conta.Pessoa
		fmt.Println(conta.Pessoa.Privilegio)
		fmt.Println(conta.Pessoa.Nome)
		switch conta.Pessoa.Privilegio {
		case 2:
			aluno, _ := sess.Get("aluno").(models.Aluno)
			estagios, _ := sess.Get("estagios").([]*models.Estagio)
			this.Data["aluno"] = aluno
			this.Data["curso"] = aluno.Curso
			this.Data["estagios"] = estagios
			this.TplName = "index_aluno.html"
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
	}
	fmt.Println(this.TplName)
	this.Render()
}

type PopularController struct {
	beego.Controller
}

func (this *PopularController) Get() {
	CriarDados()
	//this.Ctx.Redirect(302, "/")
	this.Redirect("/", 302)
}

func CriarDados() {
	o := orm.NewOrm()

	for i := 0; i < 5; i++ {
		s := strconv.Itoa(i)
		o.Insert(&models.Curso{Nome: "curso" + s})
	}
	for i := 0; i < 5; i++ {
		s := strconv.Itoa(i)
		o.Insert(&models.Pessoa{Nome: "supervisorEmpresa" + s, Privilegio: 1, Email: "supervisor@" + s})
	}
	for i := 0; i < 5; i++ {
		s := strconv.Itoa(i)
		o.Insert(&models.Pessoa{Nome: "aluno" + s, Privilegio: 2, Email: "aluno@" + s})
	}
	for i := 0; i < 5; i++ {
		s := strconv.Itoa(i)
		o.Insert(&models.Pessoa{Nome: "professor" + s, Privilegio: 3, Email: "professor@" + s})
	}
	for i := 0; i < 5; i++ {
		s := strconv.Itoa(i)
		o.Insert(&models.Pessoa{Nome: "coord" + s, Privilegio: 4, Email: "coord@" + s})
	}
	for i := 0; i < 5; i++ {
		s := strconv.Itoa(i)
		o.Insert(&models.Pessoa{Nome: "admin" + s, Privilegio: 5, Email: "admin@" + s})
	}

	var pessoas []*models.Pessoa
	qs := o.QueryTable("pessoa")
	qs.All(&pessoas)

	for i := 5; i < 25; i++ {
		_ = strconv.Itoa(i)
		fmt.Println(i)
		o.Insert(&models.Conta{Pessoa: pessoas[i], Usuario: pessoas[i].Email, Senha: pessoas[i].Nome + "_"})
	}

	var contas []*models.Conta
	qs = o.QueryTable("conta")
	qs.All(&contas)

	var cursos []*models.Curso
	qs = o.QueryTable("curso")
	qs.All(&cursos)

	for i := 0; i < 5; i++ {
		o.Insert(&models.Aluno{Conta: contas[i], Curso: cursos[i%2], Ra: uint(i)})
		o.Insert(&models.Professor{Conta: contas[i+5], Curso: cursos[i%2], Seap: contas[i+10].Usuario})
		o.Insert(&models.Professor{Conta: contas[i+10], Curso: cursos[i%4], Seap: contas[i+10].Usuario})

	}

	var alunos []*models.Aluno
	qs = o.QueryTable("aluno")
	qs.All(&alunos)

	var professores []*models.Professor
	qs = o.QueryTable("professor")
	qs.All(&professores)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		o.Insert(&models.Estagio{Aluno: alunos[i], Professor: professores[i%3]})
	}

	var estagios []*models.Estagio
	qs = o.QueryTable("estagio")
	qs.All(&estagios)

	for _, p := range pessoas {
		fmt.Println(p)
	}
	for _, p := range alunos {
		fmt.Println(p)
	}
	for _, p := range professores {
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
