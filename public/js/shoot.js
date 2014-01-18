function ShootCtrl($scope, $http, $timeout) {
    $scope.init = function() {
        var shootResult = localStorage['shootResult'];
        if (shootResult === null || shootResult == '') {
            shootResult = 'Válassz egy mezőt!';
        }
        $scope.shootResult = shootResult;

        $timeout(function() {
            $scope.shootResult = 'Válassz egy mezőt!';
            localStorage['shootResult'] = '';
        }, 5000);
    };

    $scope.shoot = function(rowS, colS) {
        $http({
            method : 'POST',
            url : '/shoot',
            data : $.param({col: String.fromCharCode(65 + colS), row: rowS + 1}),
            headers : {
                'Content-Type' : 'application/x-www-form-urlencoded'
            }
        }).
        success(function(data, status, headers, config) {
            var shootResult;
            if (data === '-1') {
                shootResult = "Nem te vagy soron!";
            } else {
                switch (data) {
                    case 'hit': shootResult = 'Talált!'; break;
                    case 'miss' : shootResult = 'Nem talált!'; break;
                    case 'hitnsunk' : shootResult = 'Talált, süllyedt!'; break;
                    case 'invalid' : shootResult = 'Válassz másik mezőt!'; break;
                }
            }
            localStorage['shootResult'] = shootResult;
            window.location.reload();
        }).
        error(function(data, status, headers, config) {
            $scope.shootResult = 'Hálózati hiba történt!';
        });
    };
}