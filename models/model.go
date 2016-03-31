package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Conta struct {
	Id      uint
	Pessoa  *Pessoa `orm:"rel(one)"`
	Usuario string
	Senha   string
}

type Pessoa struct {
	Id         uint
	Conta      *Conta `orm:"null;reverse(one)"`
	Nome       string
	Telefone   string
	Email      string `orm:"unique"`
	Privilegio int    // 0-supervisor/anonimo 1-aluno 2-professor 3-coord-curso 4-admin
}

func (s Pessoa) String() string {
	return fmt.Sprintf("(pessoa .Id: %d, .Conta: (conta), .Nome: %s, .Telefone: %s, .Email: %s, .Privilegio: %d)", s.Id, s.Nome, s.Telefone, s.Email, s.Privilegio)
}

type Aluno struct {
	Id           uint
	Conta        *Conta `orm:"rel(fk)"`
	Ra           uint   `orm:"unique"`
	Curso        *Curso `orm:"rel(fk)"`
	Periodo      uint
	CargaHoraria uint
}

func (s Aluno) String() string {
	return fmt.Sprintf("(aluno .Id: %d, .Conta: (conta), .Ra: %d, .Curso: (curso), .Periodo: %d, .CargaHoraria: %d)", s.Id, s.Ra, s.Periodo, s.CargaHoraria)
}

type Professor struct {
	Id    uint
	Conta *Conta `orm:"rel(fk)"`
	Curso *Curso `orm:"rel(fk)"`
	Seap  string `orm:"unique"`
}

func (s Professor) String() string {
	return fmt.Sprintf("(professor .Id: %d, .Conta: (conta), .Curso: (curso), .Seap: %s)", s.Id, s.Seap)
}

type Documento struct {
	Id     uint
	Titulo string
}

func (s Documento) String() string {
	return fmt.Sprintf("(documento .Id: %d, .Titulo: %s)", s.Id, s.Titulo)
}

type Estagio struct {
	Id              uint
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
	Id        uint
	Estagio   *Estagio   `orm:"rel(fk)"`
	Documento *Documento `orm:"rel(fk)"`
	//Arquivo *FILE
	DataEntrega string
}

func (s Estagio_Documento) String() string {
	return fmt.Sprintf("(estagio_documento .Id: %d, .Estagio: (estagio), .Documento: (documento), .DataEntrega: %s)", s.Id, s.DataEntrega)
}

type Curso struct {
	Id                  uint
	Nome                string `orm:"unique"`
	CHObrigatoria       uint
	CHNObrigatoria      uint
	PeriodoNObrigatorio uint
	PeriodoObrigatorio  uint
}

func (s Curso) String() string {
	return fmt.Sprintf("(curso .Id: %d, .Nome: %s, .CHObrigatoria: %d, .CHNObrigatoria: %d, .PeriodoNObrigatorio: %d, .PeriodoObrigatorio: %d)", s.Id, s.Nome, s.CHObrigatoria, s.CHNObrigatoria, s.PeriodoNObrigatorio, s.PeriodoObrigatorio)
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:estagio123@/my_db?charset=utf8", 30)
	orm.RegisterModel(new(Conta), new(Pessoa), new(Curso), new(Aluno), new(Professor), new(Estagio), new(Documento), new(Estagio_Documento))
	orm.RunSyncdb("default", true, true)
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
	qs := o.QueryTable("aluno")
	_, err = qs.All(&alunos)
	return
}
