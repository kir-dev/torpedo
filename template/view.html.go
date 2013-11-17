<html>
<head>
    <title>Megjelenítő</title>
    <link rel="stylesheet" type="text/css" href="/public/css/site.css">
</head>
<body>
    {{if hasWinner .Winner}}
        <a href="/games">Vissza az archívumhoz</a> <br/>
    {{end}}
    <div id="player-list">
        <p>Játékosok:</p>
        <ul>
            {{with .Players}}
                {{range .}}
                    <li>{{.Name}}</li>
                {{end}}
            {{end}}
        </ul>
    </div>


    <div class="winner hidden">
        A nyertes: <span>[name]</span>
    </div>

    <div id="time">
        A körből hátralevő idő: <span id="elapsed-time">0</span> s
    </div>

    <div id="players">
        <p>Jelenlegi játékos: <span id="current-player"></span></p>
        <p>Következő játékos: <span id="next-player"></span></p>
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