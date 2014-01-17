function ShootCtrl($scope, $http) {
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
            if (data === '-1') {
                $scope.shootResult = "Nem te vagy soron!";
            } else {
                $scope.shootResult = data
            }
        }).
        error(function(data, status, headers, config) {
            // TODO
            console.log("error");
        });
    }
}