const MSG_HITRESULT = 0,
      MSG_GAMESTARTED = 1,
      MSG_GAMEOVER = 2,
      MSG_ELAPSEDTIME = 3,
      MSG_TURNSTART = 4;

$(function () {

    if (window.WebSocket === undefined) {
        $("div.error").html("Your browser does not support websockets. Sorry...");
    } else {
        window.socket = initSocket();
    }

    function initSocket() {
        var url = "ws://" + window.location.host + "/ws",
            ws = new WebSocket(url);
        ws.onopen = function () {
            console.log("Socket open.");
        };
        ws.onclose = function () {
            console.log("Socket closed.");
        }
        ws.onmessage = function (event) {
            var o = JSON.parse(event.data)

            // giant switch-case to process message from server
            switch (o.type) {
                case MSG_HITRESULT:
                    shot(o.payload);
                    break;
                case MSG_ELAPSEDTIME:
                    time(o.payload);
                    break;
                case MSG_GAMESTARTED:
                    gameStarted();
                    break;
                case MSG_GAMEOVER:
                    gameEnded(o.payload);
                    break;
                case MSG_TURNSTART:
                    turn(o.payload);
                    break;
                default:
                    console.log(o);
            }
        };

        return ws;
    }

    function shot(payload) {
        $(".coord-" + payload.row + "-" + payload.col).addClass(payload.result);
    }

    function time(elapsed) {
        // TODO: get value from the server
        var remaining = 30 - Math.floor(elapsed);
        $("#elapsed-time").html(remaining);
    }

    function gameStarted() {
        console.log("Game started");
        // TODO: reset view
    }

    function gameEnded(winner) {
        $(".winner span").html(winner).parent().show();
    }

    function turn(players) {
        $("#current-player").html(players.current);
        $("#next-player").html(players.next);
    }

});