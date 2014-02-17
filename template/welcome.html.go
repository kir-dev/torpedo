<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>Kir-Dev multi torpedó</title>
    <meta name="description" content="Kir-Dev multi torpedó">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
    <meta name="HandheldFriendly" content="True">
    <meta name="MobileOptimized" content="320">

    <link href='http://fonts.googleapis.com/css?family=Open+Sans:400,300&subset=latin,latin-ext' rel='stylesheet' type='text/css'>
    <link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.0.3/css/bootstrap.min.css">
    <link rel="stylesheet" type="text/css" href="/public/css/site.css">
</head>
<body>

    <div class="row text-center">
        <div class="span12">
            <div id="header_decorator"></div>
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
        </div>
    </div>

</body>
</html>