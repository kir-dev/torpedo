<html>
    <head>
        <title>Kir-Dev multi torpedó</title>
    </head>
    <body>
        <h1>Torpedó játék</h1>

        <form action="/join" method="POST">
            <div><label for="username">Név</label></div>
            <div><input type="text" id="username" name="username" /></div>
            <div>
                <label for="is_robot">
                    <input type="checkbox" name="is_robot" id="is_robot" /> Robot vagyok
                </label>
            </div>
            <div><input type="submit" value="Csatlakozom" /></div>
        </form>
    </body>
</html>