<html>
<head>
    <title>Megjelenítő</title>
    <link rel="stylesheet" type="text/css" href="/public/css/site.css">
</head>
<body>
    {{with .Players}}
        <div class="players">
            <p>Aktuális játékosok:</p>
            <ul>
                {{range .}}
                <li>{{.Name}}</li>
                {{end}}
            </ul>
        </div>
    {{end}}

    <div class="error">
    </div>

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

    <script type="text/javascript" src="//ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js"></script>
    <script type="text/javascript" src="/public/js/board.js"></script>
</body>
</html>