<html>
<head>
    <title>Megjelenítő</title>
    <link href='http://fonts.googleapis.com/css?family=Audiowide|Open+Sans:400,300&subset=latin,latin-ext' rel='stylesheet' type='text/css'>
    <link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.0.3/css/bootstrap.min.css">
    <link rel="stylesheet" type="text/css" href="/public/css/site.css">
</head>
<body>
    <div id="header_decorator"></div>

    <div class="container">
        <div class="row">
            <div class="col-md-8">
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
            </div>
            <div class="col-md-4">
                    <img class="kirdevlogo" src="public/img/logo.png"/>
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

                        <div class="winner" style="display:none">
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

                        <h1>Csatlakozás:</h1>
                        <h2>http://stewie.sch.bme.hu:6060</h2>
                    </div>
            </div>
        </div>
    </div>



    <script type="text/javascript" src="//ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js"></script>
    <script type="text/javascript" src="/public/js/board.js"></script>
</body>
</html>