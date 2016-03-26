/*
 Copyright 2011 The Go Authors.  All rights reserved.
 Use of this source code is governed by a BSD-style
 license that can be found in the LICENSE file.
*/

function PessoasCtrl($scope, $http) {
  $scope.pessoas = [];
  $scope.working = false;

  var logError = function(data, status) {
    console.log('code '+status+': '+data);
    $scope.working = false;
  };

  var refresh = function() {
    return $http.get('/pessoas/').
      success(function(data) { $scope.pessoas = data.Pessoas; }).
      error(logError);
  };

  $scope.addPessoa = function() {
    $scope.working = true;

    $http.post('/pessoas/', {Nome: $scope.nomeText, Telefone: $scope.telefoneText, Email: $scope.emailText}).
      error(logError).
      success(function() {
        refresh().then(function() {
          $scope.working = false;
          $scope.nomeText = '';
          $scope.telefoneText = '';
          $scope.emailText = '';
        })
      });
  };

  refresh().then(function() { $scope.working = false; });
}
