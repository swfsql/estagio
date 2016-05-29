
/*
 Copyright 2011 The Go Authors.  All rights reserved.
 Use of this source code is governed by a BSD-style
 license that can be found in the LICENSE file.
*/


(function(){
  var app = angular.module('admin', [ ]);  

   app.controller('AdminCtrl', function($scope, $http, $window) { 
   this.pessoas = indv;


      /*$scope.Procura = function() {
        if($scope.pesqText == '123')
          alert("felipaaae");
        return this.tab == textRecv;
      };*/

      $scope.Procura = function(pessoa){
        return ($scope.pesqText === pessoa.idade);
      };

    });

       var indv = [
      {
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