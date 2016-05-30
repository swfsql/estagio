
/*
 Copyright 2011 The Go Authors.  All rights reserved.
 Use of this source code is governed by a BSD-style
 license that can be found in the LICENSE file.
*/


(function(){
  var app = angular.module('admin', [ ]);  

   app.controller('AdminCtrl', function($scope, $http, $window) { 
   this.pessoas = indvs;
   this.pessoa = {};

    /*this.pessoas = this.pessoas.concat(this.individuo);
    this.pessoas.push(this.individuo)*/

      this.newAluno = function (){
          alert("felipaaae0");
          this.pessoas.push(this.pessoa);
          alert("felipaaae3");

          this.pessoa = {};
      };

      this.Procura = function (pessoa){
        return ($scope.pesqText === pessoa.idade);
      };

    });

      var indv = [{nome:"jo", idade: 26}];
      var indvs = [{
          nome:"felipe",
          idade: 22,
        },
        {
          nome: "aosdsaio",
          idade: 12,
        },
        {
          nome:"felipe",
          idade: 145,
        },
        {
          nome: "fel",
          idade: 1,
        },
        {
          nome:"aes",
          idade: 1233,
        },
        {
          nome: "fas",
          idade: 3456,
        },
        {
          nome:"qweqw",
          idade: 12,
        },
        {
          nome: "aosdadewqesaio",
          idade: 42,
        },
      ];

})();