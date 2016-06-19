package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func CriarDados() {
	o := orm.NewOrm()

	for i := 0; i < 5; i++ {
		s := strconv.Itoa(i)
		o.Insert(&Curso{Nome: "curso" + s, Sigla: "sigla" + s})
	}
	for i := 0; i < 5; i++ {
		s := strconv.Itoa(i)
		o.Insert(&Pessoa{Nome: "supervisorEmpresa" + s, Privilegio: 1, Email: "supervisor@" + s + ".com"})
	}
	for i := 0; i < 5; i++ {
		s := strconv.Itoa(i)
		o.Insert(&Pessoa{Nome: "aluno" + s, Privilegio: 2, Email: "aluno@" + s + ".com"})
	}
	for i := 0; i < 5; i++ {
		s := strconv.Itoa(i)
		o.Insert(&Pessoa{Nome: "professor" + s, Privilegio: 3, Email: "professor@" + s + ".com"})
	}
	for i := 0; i < 5; i++ {
		s := strconv.Itoa(i)
		o.Insert(&Pessoa{Nome: "coord" + s, Privilegio: 4, Email: "coord@" + s + ".com"})
	}
	for i := 0; i < 5; i++ {
		s := strconv.Itoa(i)
		o.Insert(&Pessoa{Nome: "admin" + s, Privilegio: 5, Email: "admin@" + s + ".com"})
	}

	var pessoas []*Pessoa
	qs := o.QueryTable("pessoa")
	qs.All(&pessoas)

	for i := 5; i < 25; i++ {
		_ = strconv.Itoa(i)
		fmt.Println(i)
		o.Insert(&Conta{Pessoa: pessoas[i], Usuario: pessoas[i].Email, Senha: pessoas[i].Nome + "_"})
	}

	var contas []*Conta
	qs = o.QueryTable("conta")
	qs.All(&contas)

	var cursos []*Curso
	qs = o.QueryTable("curso")
	qs.All(&cursos)

	for i := 0; i < 5; i++ {
		o.Insert(&Aluno{Conta: contas[i], Curso: cursos[i%2], Ra: uint64(i)})
		o.Insert(&Professor{Conta: contas[i+5], Curso: cursos[i%2], Seap: contas[i+10].Usuario})
		o.Insert(&Professor{Conta: contas[i+10], Curso: cursos[i%4], Seap: contas[i+10].Usuario})

	}

	var alunos []*Aluno
	qs = o.QueryTable("aluno")
	qs.All(&alunos)

	var professores []*Professor
	qs = o.QueryTable("professor")
	qs.All(&professores)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		o.Insert(&Estagio{Aluno: alunos[i%5], Professor: professores[i%3]})
	}

	var estagios []*Estagio
	qs = o.QueryTable("estagio")
	qs.All(&estagios)

	for _, p := range pessoas {
		fmt.Println(p)
	}
	for _, p := range alunos {
		fmt.Println(p)
	}
	for _, p := range professores {
		fmt.Println(p)
	}
	for _, p := range estagios {
		fmt.Println(p)
	}

	//p := Pessoa{Id: pessoa.Id}
	//o.Read(&p)
	//fmt.Printf("pessoa lida: ")
	//fmt.Println(p)

}
