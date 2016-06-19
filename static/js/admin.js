
/*
   Copyright 2011 The Go Authors.  All rights reserved.
   Use of this source code is governed by a BSD-style
   license that can be found in the LICENSE file.
   */


(function(){
  var app = angular.module('admin', [ ]);  

  app.controller('AdminCtrl', function($scope, $http, $window) { 
    $scope.professores = [];   
    $scope.alunos = [];
    $scope.pessoa = {};


    this.newAluno = function (){
      var d = this.pessoa
      $http.post('/nucleo/cadastro/discente', {Ra: d.ra, Nome: d.nome, Email: d.email, Curso: d.curso, Telefone: d.telefone, Periodo: d.periodo, CargaHoraria: d.ch}).
        success(function(resp) {
          switch (resp.St) {
            case "ok":
              var curso = {Nome: d.curso}
              var pessoa = {Nome: d.nome, Email: d.email, Telefone: d.telefone, Curso: curso}
              var conta = {Pessoa: pessoa}
              var aluno = {Conta: conta, Ra: d.ra};
              $scope.alunos.push(aluno);
              d = {};
              break;
            case "err_jsoninvalido":
              alert("err_jsoninvalido");
              break;
            case "err_rajaexistente":
              alert("err_rajaexistente");
              break;
            case "err_cursonaoexistente":
              alert("err_cursonaoexistente");
              break;
          }
        });
    };

    this.newProfessor = function (){
      alert("felipaaae0");
      this.professores.push(this.pessoa);
      alert("felipaaae3");
      this.pessoa = {};
    };

    this.ProcuraProfessor = function (professor){
      if($scope.pesqText > 0)
        return ($scope.pesqText === professor.Seap);
    };
    this.ProcuraAluno = function (aluno){
      if($scope.pesqText > 0)
        return ($scope.pesqText === aluno.Ra);
    };

    this.refresh = function() {
      return $http.get('/nucleo/getdados').
        success(function(data) { 
          $scope.alunos = data.Alunos; 
          $scope.professores = data.Professores;
          //console.log(this.alunos);
        })
      //error(logError);
    };
    this.refresh();
  });

})();
