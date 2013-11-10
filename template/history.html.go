<html>
<head>
    <title>Archívum</title>
    <link rel="stylesheet" type="text/css" href="/public/css/site.css">
</head>
<body>
    <ul>
        {{range $count, $game := .}}
            <li><a href="/games/{{$game.Id}}">{{$count}}. Játék</a></li>
        {{end}}
    </ul>

    <script type="text/javascript" src="//ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js"></script>
    <script type="text/javascript" src="/public/js/board.js"></script>
</body>
</html>