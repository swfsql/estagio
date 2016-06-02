package routers

import (
	"github.com/astaxie/beego"
	"github.com/swfsql/estagio/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/discente/?:id", &controllers.AlunoController{})
	beego.Router("/docente/?:id", &controllers.ProfessorController{})
	beego.Router("/nucleo/", &controllers.AdminController{})
	beego.Router("/estagio/?:id", &controllers.EstagioController{})
	beego.Router("/popular", &controllers.PopularController{})
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/login", &controllers.LoginController{}, "get:Login;post:Post")
	beego.Router("/trocarsenha", &controllers.TrocaSenhaController{})

	beego.Router("/pessoas/", &controllers.ListaPessoaController{}, "get:PegarPessoas;post:CadastrarPessoa")

}
