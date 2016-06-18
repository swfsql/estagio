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

func GetAlunoByContaId(conta_id uint64) (aluno Aluno, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("aluno")
	err = qs.Filter("Conta", conta_id).RelatedSel().One(&aluno)
	return
}

func (s Aluno) String() string {
	return fmt.Sprintf("(aluno .Id: %d, .Conta: (conta), .Ra: %d, .Curso: (curso), .Periodo: %d, .CargaHoraria: %d)", s.Id, s.Ra, s.Periodo, s.CargaHoraria)
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
		alunos[i].Conta.Pessoa = pessoa;
		alunos[i].Conta.Senha = ""
		alunos[i].Conta.Usuario = ""

	}


	//qs := o.QueryTable("aluno").RelatedSel("curso").RelatedSel("conta").RelatedSel("conta__pessoa")
    //_, err = qs.All(&alunos)
	


	
	/*qs := o.QueryTable("aluno")
	num, err := qs.Values(&maps, "id", "ra","curso__nome", "curso__sigla","periodo","cargahoraria","conta__pessoa__nome", "conta__pessoa__telefone","conta__pessoa__email")
	alunos = make([]*Aluno, num)
	fmt.Println("huhuhu")
	fmt.Println(err)
	if err == nil {
		fmt.Println("hueuhebr")
		for i, m := range maps {
			//fmt.Println(m)
			//fmt.Println("varias linhas")
			//fmt.Println(reflect.TypeOf(m["id"]))
			aluno := &Aluno{}

			switch v := m["id"].(type) {
			case uint64:
				aluno.Id = v;
			} 
			switch v := m["ra"].(type) {
			case uint64:
				aluno.Ra = v;
			} 
			switch v := m["periodo"].(type) {
			case uint64:
				aluno.Periodo = v;
			} 
			switch v := m["cargahoraria"].(type) {
			case uint64:
				aluno.CargaHoraria = v;
			} 
			//aluno.Id = m["id"].(uint64);
				//Id: m["id"].(nil), Ra: m["ra"].(uint64), Periodo: m["periodo"].(uint64), CargaHoraria: m["cargahoraria"].(uint64)}
			aluno.Curso = &Curso{}
			switch v := m["curso__nome"].(type) {
			case string:
				aluno.Curso.Nome = v;
			}
			switch v := m["curso__sigla"].(type) {
			case string:
				aluno.Curso.Sigla = v;
			}
			//Nome: m["curso__nome"].(string), Sigla: m["curso__sigla"].(string)
			aluno.Conta = &Conta{}
			aluno.Conta.Pessoa = &Pessoa{}
			switch v := m["conta__pessoa__nome"].(type) {
			case string:
				aluno.Conta.Pessoa.Nome = v;
			}
			switch v := m["conta__pessoa__telefone"].(type) {
			case string:
				aluno.Conta.Pessoa.Telefone = v;
			}
			switch v := m["conta__pessoa__email"].(type) {
			case string:
				aluno.Conta.Pessoa.Email = v;
			}
			//Nome: m["conta__pessoa__nome"].(string), Telefone: m["conta__pessoa__telefone"].(string), Email: m["conta__pessoa__email"].(string)
			alunos[i] = aluno
			fmt.Println(alunos[i])
		}
	}*/

	//_, err = qs.All(&alunos)
	return
}

func GetEstagios() (estagios []*Estagio, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("estagio")
	_, err = qs.All(&estagios)
	return
}

	/*//aluno
	Id           uint64
	Conta        *Conta `orm:"rel(fk)"`
	Ra           uint64 `64orm:"unique"`
	Curso        *Curso `orm:"rel(fk)"`
	Periodo      uint64
	CargaHoraria uint64

	//pessoa
	Id         uint64
	Conta      *Conta `orm:"null;reverse(one)"`
	Nome       string
	Telefone   string `orm:"null"`
	Email      string `orm:"unique"`
	Privilegio uint32 // 0-supervisor/anonimo 1-aluno 2-professor 3-coord-curso 4-admin


type Conta struct {
	Id      uint64
	Pessoa  *Pessoa `orm:"rel(one)"`
	Usuario string
	Senha   string `orm:"null"`
}*/