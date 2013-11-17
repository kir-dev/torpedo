<html>
<head>
    <title>Archívum</title>
    <link rel="stylesheet" type="text/css" href="/public/css/site.css">
</head>
<body>
    {{if isHistoryEmpty .}}
        Nincs még lejátszott játék.
    {{else}}
        <h2>Eddigi játszmák:</h2>
        <ul>
            {{range $count, $game := .}}
                <li><a href="/games/{{$game.Id}}">{{add $count 1}}. Játék</a></li>
            {{end}}
        </ul>
    {{end}}

    <script type="text/javascript" src="//ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js"></script>
    <script type="text/javascript" src="/public/js/board.js"></script>
</body>
</html>