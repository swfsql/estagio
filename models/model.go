package models

import (
	"errors"
	"fmt"
	//"reflect"
	"github.com/astaxie/beego/orm"
)

// ERRS
var (
	ErrNoRows = errors.New("<QuerySeter> no row found")
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:estagio123@/my_db?charset=utf8", 30)
	orm.RegisterModel(new(Conta), new(Pessoa), new(Curso), new(Aluno), new(Professor), new(Estagio), new(Documento), new(Estagio_Documento))
	orm.RunSyncdb("default", true, true)
	//orm.Debug = true
	CriarDados()
}

type Conta struct {
	Id      uint64
	Pessoa  *Pessoa `orm:"rel(one)"`
	Usuario string
	Senha   string `orm:"null"`
}

func GetContaByEmail(email string) (conta Conta, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("conta")
	err = qs.Filter("Usuario", email).RelatedSel().One(&conta)
	if err == orm.ErrNoRows {
		err = ErrNoRows
	}
	return
}

type Pessoa struct {
	Id         uint64
	Conta      *Conta `orm:"null;reverse(one)"`
	Nome       string
	Telefone   string `orm:"null"`
	Email      string `orm:"unique"`
	Privilegio uint32 // 0-supervisor/anonimo 1-aluno 2-professor 3-coord-curso 4-admin
}

func (s Pessoa) String() string {
	return fmt.Sprintf("(pessoa .Id: %d, .Conta: (conta), .Nome: %s, .Telefone: %s, .Email: %s, .Privilegio: %d)", s.Id, s.Nome, s.Telefone, s.Email, s.Privilegio)
}

type Aluno struct {
	Id           uint64
	Conta        *Conta `orm:"rel(fk)"`
	Ra           uint64 `orm:"unique"`
	Curso        *Curso `orm:"rel(fk)"`
	Periodo      uint64
	CargaHoraria uint64
}

func (s Aluno) String() string {
	return fmt.Sprintf("(aluno .Id: %d, .Conta: (conta), .Ra: %d, .Curso: (curso), .Periodo: %d, .CargaHoraria: %d)", s.Id, s.Ra, s.Periodo, s.CargaHoraria)
}
func GetAlunoByContaId(conta_id uint64) (aluno Aluno, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("aluno")
	err = qs.Filter("Conta", conta_id).RelatedSel().One(&aluno)
	return
}
func GetAlunoByAlunoId(aluno_id uint64) (aluno Aluno, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("aluno")
	err = qs.Filter("Id", aluno_id).RelatedSel().One(&aluno)
	return
}
func GetAlunoByRa(ra uint64) (aluno Aluno, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("aluno")
	err = qs.Filter("Ra", ra).RelatedSel().One(&aluno)
	if err == orm.ErrNoRows {
		err = ErrNoRows
	}
	return
}
func CadastroAluno(aluno *Aluno) {
	o := orm.NewOrm()
	o.Insert(aluno.Conta.Pessoa)
	o.Insert(aluno.Conta)
	o.Insert(aluno)
	//o.Insert(&Aluno{Conta: aluno.Conta, Curso: cursos[i%2], Ra: uint64(i)})
	//	o.Insert(aluno)
}

type Professor struct {
	Id    uint64
	Conta *Conta `orm:"rel(fk)"`
	Curso *Curso `orm:"rel(fk)"`
	Seap  string `orm:"unique"`
}

func (s Professor) String() string {
	return fmt.Sprintf("(professor .Id: %d, .Conta: (conta), .Curso: (curso), .Seap: %s)", s.Id, s.Seap)
}

func GetProfessorByContaId(conta_id uint64) (professor Professor, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("professor")
	err = qs.Filter("Conta", conta_id).RelatedSel().One(&professor)
	return
}

type Documento struct {
	Id     uint64
	Titulo string
}

func (s Documento) String() string {
	return fmt.Sprintf("(documento .Id: %d, .Titulo: %s)", s.Id, s.Titulo)
}

type Estagio struct {
	Id              uint64
	Aluno           *Aluno     `orm:"rel(fk)"`
	Professor       *Professor `orm:"rel(fk)"`
	LocalFis        string
	Obrigatoriedade bool
	DataInicio      string
	DataFim         string
}

func (s Estagio) String() string {
	return fmt.Sprintf("(estagio .Id: %d, .Aluno: (aluno), .Professor: (professor), .LocalFis: %s, .Obrigatoriedade: %t, .DataInicio: %s, .DataFim: %s)", s.Id, s.LocalFis, s.Obrigatoriedade, s.DataInicio, s.DataFim)
}

type Estagio_Documento struct {
	Id        uint64
	Estagio   *Estagio   `orm:"rel(fk)"`
	Documento *Documento `orm:"rel(fk)"`
	//Arquivo *FILE
	DataEntrega string
}

func (s Estagio_Documento) String() string {
	return fmt.Sprintf("(estagio_documento .Id: %d, .Estagio: (estagio), .Documento: (documento), .DataEntrega: %s)", s.Id, s.DataEntrega)
}

type Curso struct {
	Id                  uint64
	Nome                string `orm:"unique"`
	Sigla               string `orm:"unique"`
	CHObrigatoria       uint64
	CHNObrigatoria      uint64
	PeriodoNObrigatorio uint64
	PeriodoObrigatorio  uint64
}

func (s Curso) String() string {
	return fmt.Sprintf("(curso .Id: %d, .Nome: %s, .Sigla: %s, .CHObrigatoria: %d, .CHNObrigatoria: %d, .PeriodoNObrigatorio: %d, .PeriodoObrigatorio: %d)", s.Id, s.Nome, s.Sigla, s.CHObrigatoria, s.CHNObrigatoria, s.PeriodoNObrigatorio, s.PeriodoObrigatorio)
}
func GetCursoBySigla(sigla string) (curso Curso, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("curso")
	err = qs.Filter("Sigla", sigla).RelatedSel().One(&curso)
	if err == orm.ErrNoRows {
		err = ErrNoRows
	}
	return
}
func GetCursosSiglas() (siglas []string, err error) {
	o := orm.NewOrm()
	var list []orm.ParamsList
	num, err := o.QueryTable("curso").ValuesList(&list, "sigla")
	siglas = make([]string, num)
	if err == nil {
		for i, v := range list {
			siglas[i] = v[0].(string)
		}
	}
	return
}

func (p *Professor) GetEstagios() (estagios []*Estagio, err error) {
	//var estagios []*Estagio
	o := orm.NewOrm()
	qs := o.QueryTable("estagio")
	_, err = qs.Filter("Professor", p.Id).RelatedSel().All(&estagios)

	/*for _, estagio := range estagios {
		fmt.Printf("Id: %d, ", estagio.Id, estagio.Aluno, estagio.Professor)
	}*/
	return
}

func GetEstagiosByCurso(curso *Curso) (estagios []*Estagio, err error) {
	//var estagios []*Estagio
	o := orm.NewOrm()
	qs := o.QueryTable("estagio")
	_, err = qs.Filter("Professor__Curso__Id", curso.Id).RelatedSel().All(&estagios)
	return
}

func (a *Aluno) GetEstagios() (estagios []*Estagio, err error) {
	//var estagios []*Estagio
	o := orm.NewOrm()
	qs := o.QueryTable("estagio")
	_, err = qs.Filter("Aluno", a.Id).RelatedSel().All(&estagios)

	/*for _, estagio := range estagios {
		fmt.Printf("Id: %d, ", estagio.Id, estagio.Aluno, estagio.Professor)
	}*/
	return
}

func GetProfessores() (professores []*Professor, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("professor")
	_, err = qs.All(&professores)
	return
}

func GetAlunos() (alunos []*Aluno, err error) {
	o := orm.NewOrm()

	qs := o.QueryTable("aluno").RelatedSel("curso").RelatedSel("conta")
	_, err = qs.All(&alunos)

	for i, a := range alunos {
		pessoa := &Pessoa{}
		o.Read(pessoa)
		err = o.QueryTable("pessoa").Filter("id", a.Conta.Pessoa.Id).Limit(1).One(pessoa)
		alunos[i].Conta.Pessoa = pessoa
		alunos[i].Conta.Senha = ""
		alunos[i].Conta.Usuario = ""

	}

	fmt.Println("alunos:")
	fmt.Println(alunos)

	return
}

func GetEstagios() (estagios []*Estagio, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("estagio")
	_, err = qs.All(&estagios)
	return
}
