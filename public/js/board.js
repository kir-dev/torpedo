const MSG_HITRESULT = 0,
      MSG_GAMESTARTED = 1,
      MSG_GAMEOVER = 2,
      MSG_ELAPSEDTIME = 3;

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

            switch (o.type) {
                case MSG_HITRESULT:
                    shotReceived(o.payload);
                    break;
                default:
                    console.log(o);
            }
        };

        return ws;
    }

    function shotReceived(payload) {
        $(".coord-" + payload.row + "-" + payload.col).addClass(payload.result);
    }
});