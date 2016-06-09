
/*
 Copyright 2011 The Go Authors.  All rights reserved.
 Use of this source code is governed by a BSD-style
 license that can be found in the LICENSE file.
*/


(function(){
  var app = angular.module('admin', [ ]);  

   app.controller('AdminCtrl', function($scope, $http, $window) { 
   this.professores = indvs;   
   this.alunos = indvs;
   this.pessoa = {};


      this.newAluno = function (){
          alert("felipaaae0");
          this.alunos.push(this.pessoa);
          alert("felipaaae3");
          this.pessoa = {};
      };

      this.newProfessor = function (){
          alert("felipaaae0");
          this.professores.push(this.pessoa);
          alert("felipaaae3");
          this.pessoa = {};
      };

      this.ProcuraProfessor = function (pessoa){
        return ($scope.pesqText === pessoa.seap);
      };
      this.ProcuraAluno = function (pessoa){
        return ($scope.pesqText === pessoa.ra);
      };

    });

      var indvs = [{
          nome:"felipe",
          ra: 22,
        },
        {
          nome: "aosdsaio",
          ra: 12,
        },
        {
          nome:"felipe123",
          seap: 999,
        },
        {
          nome: "fel",
          ra: 1,
        },
        {
          nome:"aes",
          ra: 1233,
        },
        {
          nome: "fas",
          ra: 3456,
        },
        {
          nome:"qweqw",
          ra: 12,
        },
        {
          nome: "aosdadewqesaio",
          ra: 42,
        },];

})();