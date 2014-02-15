<html>
<head>
    <title>Megjelenítő</title>
    <link href='http://fonts.googleapis.com/css?family=Audiowide|Open+Sans:400,300&subset=latin,latin-ext' rel='stylesheet' type='text/css'>
    <link rel="stylesheet" type="text/css" href="/public/css/site.css">
</head>
<body>
    <div id="header_decorator"></div>

    <div id="game-info">
        {{if hasWinner .Winner}}
            <a href="/games">Vissza az archívumhoz</a> <br/>
        {{end}}
        <div id="player-list">
            <p class="title">Játékosok:</p>
            <ul>
                {{with .Players}}
                    {{range .}}
                        <li>
                            {{.Name}}
                            (talált: <span class="player-color-info" style="background-color: {{.Color.Hit}};"></span>,
                            talált &amp; süllyedt: <span class="player-color-info" style="background-color: {{.Color.HitAndSunk}};"></span>)

                        </li>
                    {{end}}
                {{end}}
            </ul>
            <div class="clearfix"></div>
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
            <div class="clearfix"></div>
        </div>
    </div>

    <table id="board">
        {{with .Board}}
            {{letters (len .Fields) "td"}}

            {{range $col, $ := .Fields}}
                <tr>
                    <td>{{add $col 1}}</td>
                    {{range $row, $ := .}}
                        <td class="field coord-{{$col}}-{{$row}}" style="background-color: {{ship_color .}};"></td>
                    {{end}}
                </tr>
            {{end}}
        {{end}}
    </table>

    <script type="text/javascript" src="//ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js"></script>
    <script type="text/javascript" src="/public/js/board.js"></script>
</body>
</html>