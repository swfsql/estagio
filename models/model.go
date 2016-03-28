package models

import (
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

type Aluno struct {
	Id           uint
	Conta        *Conta `orm:"rel(fk)"`
	Ra           uint   `orm:"unique"`
	Curso        *Curso `orm:"rel(fk)"`
	Periodo      uint
	CargaHoraria uint
}

type Professor struct {
	Id    uint
	Conta *Conta `orm:"rel(fk)"`
	Curso *Curso `orm:"rel(fk)"`
	Seap  string `orm:"unique"`
}

type Documento struct {
	Id     uint
	Titulo string
}

type Coord struct {
	Id    uint
	Conta *Conta `orm:"rel(fk)"`
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

type Estagio_Documento struct {
	Id        uint
	Estagio   *Estagio   `orm:"rel(fk)"`
	Documento *Documento `orm:"rel(fk)"`
	//Arquivo *FILE
	DataEntrega string
}

type Curso struct {
	Id                  uint
	Nome                string `orm:"unique"`
	CHObrigatoria       uint
	CHNObrigatoria      uint
	PeriodoNObrigatorio uint
	PeriodoObrigatorio  uint
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:estagio123@/my_db?charset=utf8", 30)
	orm.RegisterModel(new(Conta), new(Pessoa), new(Curso), new(Aluno), new(Professor), new(Coord), new(Estagio), new(Documento), new(Estagio_Documento))
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
