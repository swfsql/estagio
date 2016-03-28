package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/swfsql/estagio/models"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
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

	//md5senha := md5.New()
	//io.WriteString(md5senha, dado.Senha)
	//buffer := bytes.NewBuffer(nil)
	//fmt.Fprintf(buffer, "%x", md5senha.Sum(nil))
	//dado.Senha = buffer.String()

	//o := orm.NewOrm()
	erro := struct{ erro string }{""}
	//var conta models.Conta = models.Conta{Usuario: dado.Email}
	var conta models.Conta = models.Conta{}
	o := orm.NewOrm()
	o.QueryTable("conta").Filter("Usuario", dado.Email).RelatedSel().One(&conta)
	//err = o.Read(&conta, "Usuario")
	if err == orm.ErrNoRows {
		erro.erro = "Usuario inexistente!"
		fmt.Println("Usuario nao cadastrado!")
	} else if dado.Senha != conta.Senha {
		fmt.Println("usuario nao logado!")
		fmt.Printf("%s nao bate com %s!\n", dado.Senha, conta.Senha)
		erro.erro = "algo deu errado"
	} else {
		//Set the session successful login
		sess := this.StartSession()
		//sess := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)

		sess.Set("conta", conta)
		fmt.Println(conta.Pessoa.Privilegio)
		fmt.Println(conta.Pessoa.Nome)
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

			break
		}
		fmt.Println("usuario logado!")
		this.Ctx.Redirect(302, "/")
		return
	}

	// erro
	this.Data["json"] = erro
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
