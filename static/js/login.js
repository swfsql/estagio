/*
 Copyright 2011 The Go Authors.  All rights reserved.
 Use of this source code is governed by a BSD-style
 license that can be found in the LICENSE file.
*/

function LoginCtrl($scope, $http, $window) {
  $scope.working = false;

  var logError = function(data, status) {
    console.log('code '+status+': '+data);
    $scope.working = false;
  };

  $scope.Entrar = function() {
    $scope.working = true;

    $http.post('/login', {Email: $scope.emailText, Senha: $scope.senhaText}).
      error(logError).
      success(function(st) {
          alert("aehoo");
          $scope.st = st.Status;
          $scope.working = false;

          alert(st.Status);
          
          switch (st.Status) {
           case "ok":
              alert("ok blabalbal");
              $window.location.href = '/';
            break;
           case "err_usuario_inexiste":
              alert("err_usuario_inexiste");

            break;
           case "err_senha_invalida":
              alert("err_senha_invalida");

            break;

          }

          //$scope.emailText = '';
          //$scope.senhaText = '';
        
      });
  };

}

