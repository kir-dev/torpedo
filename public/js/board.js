const MSG_HITRESULT = 0,
      MSG_GAMESTARTED = 1,
      MSG_GAMEOVER = 2,
      MSG_ELAPSEDTIME = 3,
      MSG_TURNSTART = 4,
      MSG_PLAYERJOINED = 5;

$(function () {

    if (window.WebSocket === undefined) {
        $("div.error").html("Your browser does not support websockets. Sorry...");
    } else {
        window.socket = initSocket();
    }

    function initSocket() {
        var url = "ws://" + window.location.host + "/ws",
            ws = new WebSocket(url),
            events = {};

        events[MSG_HITRESULT] = shot;
        events[MSG_ELAPSEDTIME] = time;
        events[MSG_GAMESTARTED] = gameStarted;
        events[MSG_GAMEOVER] = gameEnded;
        events[MSG_TURNSTART] = turn;
        events[MSG_PLAYERJOINED] = joined;

        ws.onopen = function () {
            console.log("Socket open.");
        };
        ws.onclose = function () {
            console.log("Socket closed.");
        };
        ws.onmessage = function (event) {
            var o = JSON.parse(event.data),
                f = events[o.type];

            if (f !== undefined) {
                f(o.payload);
            } else {
                console.log(o);
            }
        };

        return ws;
    }

    function shot(payload) {
        $(".coord-" + payload.row + "-" + payload.col).attr("style", "background-color:"+payload.result+";")
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

    function joined(player) {
        /*$("#player-list ul").append("<li>" + player + "</li>");*/
        window.location.reload();
    }
});