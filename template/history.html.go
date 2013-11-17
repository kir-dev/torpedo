<html>
<head>
    <title>Archívum</title>
    <link rel="stylesheet" type="text/css" href="/public/css/site.css">
</head>
<body>
    <h2>Eddigi játszmák:</h2>
    <ul>
        {{range $count, $game := .}}
            <li><a href="/games/{{$game.Id}}">{{add $count 1}}. Játék</a></li>
        {{else}}
            Nincs még lejátszott játék.
        {{end}}
    </ul>

    <script type="text/javascript" src="//ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js"></script>
    <script type="text/javascript" src="/public/js/board.js"></script>
</body>
</html>