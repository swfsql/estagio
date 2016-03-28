/*
 Copyright 2011 The Go Authors.  All rights reserved.
 Use of this source code is governed by a BSD-style
 license that can be found in the LICENSE file.
*/

function LoginCtrl($scope, $http) {
  $scope.working = false;

  var logError = function(data, status) {
    console.log('code '+status+': '+data);
    $scope.working = false;
  };

  $scope.Entrar = function() {
    $scope.working = true;

    $http.post('/login', {Email: $scope.emailText, Senha: $scope.senhaText}).
      error(logError).
      success(function() {
        
          $scope.working = false;
          $scope.emailText = '';
          $scope.senhaText = '';
        
      });
  };

}

