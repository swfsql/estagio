package models

import (
	"github.com/astaxie/beego/orm"
)

type Pessoa struct {
	Id    uint
	Nome  string
	Email string
}

type Curso struct {
	Id   uint
	Nome string
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
	Id     uint
	Pessoa *Pessoa `orm:"rel(fk)"`
	Curso  *Curso  `orm:"rel(fk)"`
	Codigo uint
}

type Coord struct {
	Id     uint
	Pessoa *Pessoa `orm:"rel(fk)"`
}

type Estagio struct {
	Id    uint
	Aluno *Aluno `orm:"rel(fk)"`
	Prof  *Prof  `orm:"rel(fk)"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:estagio123@/my_db?charset=utf8", 30)
	orm.RegisterModel(new(Pessoa), new(Curso), new(Aluno), new(Prof), new(Coord), new(Estagio))
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
