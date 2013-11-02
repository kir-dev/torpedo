<!DOCTYPE html>
<html>
<head>
    <title>HIBA - {{.Title}}</title>
</head>
<body>
    <h1>Hiba történt: {{.Title}}</h1>
    {{if .IsDev}}
        <h3>Cause</h3>
        <p class="error-msg">
            {{.Message}}
        </p>
    {{end}}
</body>
</html>