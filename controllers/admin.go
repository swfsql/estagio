package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/swfsql/estagio/models"
)

// ERRS
var (
	st_err_jsoninvalido      string = "err_jsoninvalido"
	st_err_rajaexistente     string = "err_rajaexistente"
	st_err_cursonaoexistente string = "err_cursonaoexistente"
)

type AdminController struct {
	AuthAdminController
}

func (this *AdminController) Nucleo() {
	sess := this.StartSession()
	conta := sess.Get("conta").(models.Conta)
	fmt.Println("ENTROU NA PAGINA DO Admin")
	this.Data["Usuario"] = conta.Pessoa
	this.Data["CursosSiglas"], _ = models.GetCursosSiglas()
	fmt.Println(conta.Pessoa.Privilegio)
	fmt.Println(conta.Pessoa.Nome)
	this.TplName = "admin.html"
	this.All()
	this.Render()
}

func (this *AdminController) AdiquirirDados() {
	alunos, _ := models.GetAlunos()
	professores, _ := models.GetProfessores()
	dado := struct {
		Alunos      []*models.Aluno
		Professores []*models.Professor
	}{alunos, professores}
	this.Data["json"] = dado
	this.ServeJSON()
}

func (this *AdminController) CadastroDiscente() {
	dado := struct {
		Ra           uint64
		Nome         string
		Email        string
		Curso        string
		Telefone     string
		Periodo      uint64
		CargaHoraria uint64
	}{}
	wjson := struct {
		St string
	}{""}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &dado)
	if err != nil {
		fmt.Println(err)
		//this.Ctx.Output.SetStatus(400)
		//this.Ctx.Output.Body([]byte("JSON invalido"))
		wjson.St = st_err_jsoninvalido
		this.Data["json"] = wjson
		this.ServeJSON()
		return
	}

	fmt.Println(dado)

	_, err_noRow := models.GetAlunoByRa(dado.Ra)
	if err_noRow == nil {
		fmt.Println(st_err_rajaexistente)
		wjson.St = st_err_rajaexistente
		this.Data["json"] = wjson
		this.ServeJSON()
		return
	}

	var curso models.Curso
	curso, err_noRow = models.GetCursoBySigla(dado.Curso)
	if err_noRow != nil {
		fmt.Println(st_err_cursonaoexistente)
		wjson.St = st_err_cursonaoexistente
		this.Data["json"] = wjson
		this.ServeJSON()
		return
	}

	pessoa := &models.Pessoa{Conta: nil, Nome: dado.Nome, Telefone: dado.Telefone, Email: dado.Email, Privilegio: 2}
	conta := &models.Conta{Pessoa: pessoa, Usuario: dado.Email, Senha: dado.Nome + "_"}
	pessoa.Conta = conta
	aluno := &models.Aluno{Conta: conta, Ra: dado.Ra, Curso: &curso, Periodo: dado.Periodo, CargaHoraria: dado.CargaHoraria}

	models.CadastroAluno(aluno)

	fmt.Println(st_ok)
	wjson.St = st_ok
	this.Data["json"] = wjson
	this.ServeJSON()
	return
}
