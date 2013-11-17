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

    <form action="/shoot" method="POST">
        <input type="text" name="col" placeholder="A" size="1"/>
        <input type="text" name="row" placeholder="1" size="1"/>
        <input type="submit" value="Tüzel" />
    </form>
    {{with .Feedback}}
        <p class="hit-result">
            {{.}}
        </p>
    {{end}}

    {{with .Game}}
        <table id="board">
            {{with .Board}}
                {{letters (len .Fields) "td"}}

                {{range $col, $ := .Fields}}
                    <tr>
                        <td>{{add $col 1}}</td>
                        {{range $row, $ := .}}
                            <td class="field coord-{{$col}}-{{$row}} {{ship_class .}}"></td>
                        {{end}}
                    </tr>
                {{end}}
            {{end}}
        </table>
    {{end}}

    <script type="text/javascript" src="//ajax.googleapis.com/ajax/libs/angularjs/1.2.1/angular.min.js"></script>
    <script type="text/javascript" src="/public/js/shoot.js"></script>
</body>
</html>