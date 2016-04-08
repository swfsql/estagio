package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/swfsql/estagio/models"
)

var (
	st_ok                   string = "ok"
	st_err_usuario_inexiste string = "err_usuario_inexiste"
	st_err_senha_invalida   string = "err_senha_invalida"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
	this.Data["HeadTitle"] = "Login Title"
	this.Data["HeadScripts"] = []string{"/static/js/login.js"}
	this.Data["HeadStyles"] = []string{"/static/css/login.css"}
	this.Render()
}

func (this *LoginController) Post() {
	dado := struct {
		Email string
		Senha string
	}{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &dado)
	if err != nil {
		fmt.Println(err)
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("JSON invalido"))
		return
	}

	fmt.Println(dado)

	//md5senha := md5.New()
	//io.WriteString(md5senha, dado.Senha)
	//buffer := bytes.NewBuffer(nil)
	//fmt.Fprintf(buffer, "%x", md5senha.Sum(nil))
	//dado.Senha = buffer.String()

	status := struct{ Status string }{""}

	var conta models.Conta = models.Conta{}
	o := orm.NewOrm()
	o.QueryTable("conta").Filter("Usuario", dado.Email).RelatedSel().One(&conta)

	if err == orm.ErrNoRows {
		status.Status = st_err_usuario_inexiste
		this.Data["json"] = status
		this.ServeJSON()
		return

	}

	fmt.Println(conta)

	if dado.Senha != conta.Senha {
		fmt.Printf("%s nao bate com %s!\n", dado.Senha, conta.Senha)
		status.Status = st_err_senha_invalida
		this.Data["json"] = status
		this.ServeJSON()
		return

	}

	sess := this.StartSession()
	sess.Set("conta", conta)

	switch conta.Pessoa.Privilegio {

	case 2: // aluno
		var aluno models.Aluno
		o.QueryTable("aluno").Filter("Conta", conta.Id).RelatedSel().One(&aluno)
		estagios, _ := aluno.GetEstagios()
		sess.Set("aluno", aluno)
		sess.Set("estagios", estagios)
		break

	case 3: // professor
		var professor models.Professor
		sess.Set("professor", professor)
		// pegar estagios
		// ...
		break

	case 4: // coord curso
		// pegar estagios
		// ...
		break

	case 5: // admin
		// ...
		break
	}
	status.Status = st_ok
	this.Data["json"] = status
	this.ServeJSON()

}

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.TplName = "registrar.html"
	this.Render()
}

func (this *RegisterController) Post() {
	dado := struct {
		Nome       string
		Telefone   string
		Email      string
		Senha      string
		Privilegio string
	}{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &dado)
	if err != nil {
		fmt.Println(err)
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("JSON invalido"))
		return
	}

	//md5senha := md5.New()
	//io.WriteString(md5senha, dado.Senha)
	//buffer := bytes.NewBuffer(nil)
	//fmt.Fprintf(buffer, "%x", md5senha.Sum(nil))
	//dado.Senha = buffer.String()

	//p = &models.Pessoa{Nome: dado.Nome, Email: dado.Email, Telefone: dado.Telefone, Senha: dado.Senha, Privilegio: dado.Privilegio}

	//o := orm.NewOrm()
	//o.Insert(p)
}
