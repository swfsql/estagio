package routers

import (
	"github.com/astaxie/beego"
	"github.com/swfsql/estagio/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/popular", &controllers.PopularController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/pessoas/", &controllers.ListaPessoaController{}, "get:PegarPessoas;post:CadastrarPessoa")

}
