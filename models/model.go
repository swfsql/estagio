package models

import (
	"github.com/astaxie/beego/orm"
)

type Pessoa struct {
	Id       uint
	Nome     string
	Telefone string
	Email    string
	Cadastro string
	Senha    string
}

type Aluno struct {
	Id           uint
	Pessoa       *Pessoa `orm:"rel(fk)"`
	Ra           uint
	Curso        *Curso `orm:"rel(fk)"`
	Periodo      uint
	CargaHoraria uint
}

type Prof struct {
	Id         uint
	Pessoa     *Pessoa `orm:"rel(fk)"`
	Curso      *Curso  `orm:"rel(fk)"`
	Seap       string
	Privilegio uint
}

type Documento struct {
	Id     uint
	Titulo string
}

type Coord struct {
	Id     uint
	Pessoa *Pessoa `orm:"rel(fk)"`
}

type Estagio struct {
	Id              uint
	Aluno           *Aluno `orm:"rel(fk)"`
	Prof            *Prof  `orm:"rel(fk)"`
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
	Nome                string
	CHObrigatoria       uint
	CHNObrigatoria      uint
	PeriodoNObrigatorio uint
	PeriodoObrigatorio  uint
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:estagio123@/my_db?charset=utf8", 30)
	orm.RegisterModel(new(Pessoa), new(Curso), new(Aluno), new(Prof), new(Coord), new(Estagio), new(Documento), new(Estagio_Documento))
	orm.RunSyncdb("default", true, true)
}

func (p *Prof) GetEstagios() (estagios []*Estagio, err error) {
	//var estagios []*Estagio
	o := orm.NewOrm()
	qs := o.QueryTable("estagio")
	_, err = qs.Filter("Prof", p.Id).RelatedSel().All(&estagios)

	/*for _, estagio := range estagios {
		fmt.Printf("Id: %d, ", estagio.Id, estagio.Aluno, estagio.Prof)
	}*/
	return
}

func GetProfs() (profs []*Prof, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("prof")
	_, err = qs.All(&profs)
	return
}

func GetAlunos() (alunos []*Aluno, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("aluno")
	_, err = qs.All(&alunos)
	return
}
