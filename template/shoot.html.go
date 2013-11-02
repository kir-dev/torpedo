<html>
<head>
    <title>Irányítás</title>
</head>
<body>
    <form action="/shoot" method="POST">
        <input type="text" name="col" placeholder="A" size="1"/>
        <input type="text" name="row" placeholder="1" size="1"/>
        <input type="submit" value="Tüzel" />
    </form>
    <p class="hit-result">
        {{.}}
    </p>
</body>
</html>