
/*
 Copyright 2011 The Go Authors.  All rights reserved.
 Use of this source code is governed by a BSD-style
 license that can be found in the LICENSE file.
*/


(function(){
  var app = angular.module('admin', [ ]);  

   app.controller('AdminCtrl', function($scope, $http, $window) { 
   this.pessoas = indv;
   this.tab == '123'; 

   alert("felsdfgsdipe");

      $scope.Procura = function() {
        if($scope.pesqText == '123')
          alert("felipaaae");
        return this.tab == textRecv;
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
      }
    ];

})();