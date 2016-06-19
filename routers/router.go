package routers

import (
	"github.com/astaxie/beego"
	"github.com/swfsql/estagio/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/discente/?:id", &controllers.AlunoController{})

	beego.Router("/docente/?:id", &controllers.ProfessorController{})

	beego.Router("/nucleo/", &controllers.AdminController{}, "get:Nucleo")
	beego.Router("/nucleo/getdados", &controllers.AdminController{}, "get:AdiquirirDados")
	beego.Router("/nucleo/cadastro/discente", &controllers.AdminController{}, "post:CadastroDiscente")

	beego.Router("/estagio/?:id", &controllers.EstagioController{})

	beego.Router("/login", &controllers.LoginController{}, "get:Login;post:Post")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/trocarsenha", &controllers.TrocaSenhaController{})
}
