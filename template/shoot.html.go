<!DOCTYPE html>
<html ng-app>
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>Irányítás</title>
    <meta name="description" content="Torpedó irányítás">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
    <meta name="HandheldFriendly" content="True">
    <meta name="MobileOptimized" content="320">

    <link rel="stylesheet" type="text/css" href="/public/css/site.css">
</head>
<body ng-controller="ShootCtrl">

    <input ng-model="shootResult" />
    <input type="button" value="Jelenlegi állás" onclick="window.location.reload();" />

    {{with .Game}}
        <table id="board">
            {{with .Board}}
                {{letters (len .Fields) "td"}}

                {{range $col, $ := .Fields}}
                    <tr>
                        <td>{{add $col 1}}</td>
                        {{range $row, $ := .}}
                            <td class="field coord-{{$col}}-{{$row}} {{ship_class .}}" ng-click="shoot({{$col}},{{$row}})"></td>
                        {{end}}
                    </tr>
                {{end}}
            {{end}}
        </table>
    {{end}}

    <script type="text/javascript" src="//ajax.googleapis.com/ajax/libs/angularjs/1.2.1/angular.min.js"></script>
    <script type="text/javascript" src="//code.jquery.com/jquery-1.10.1.min.js"></script>
    <script type="text/javascript" src="/public/js/shoot.js"></script>
</body>
</html>