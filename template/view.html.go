<html>
<head>
    <title>Megjelenítő</title>
</head>
<body>
    <ul>
        {{range .}}
            <li>{{.Name}} and he is {{if not .IsBot}}not{{end}} a bot.</li>
        {{end}}
    </ul>
</body>
</html>