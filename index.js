// read dom elements
var results = document.getElementById("results");
var form = document.getElementById("form");
var input = document.getElementById("input");

var websocket = null;
var websocketReady = false;
var websocketSupported = "WebSocket" in window;

(function(){
    if(!websocketSupported) {
        logError("websocket is not supported!");
        return;
    }

    // connect to the websocket on the current host, using 'wss://' if on https, else 'ws://'
    var websocketAddr = (location.protocol == "https" ? "wss://" : "ws://") + location.host + "/socket";
    websocket = new WebSocket(websocketAddr);

    websocket.onopen = function() {
        websocketReady = true;
        logInfo("websocket connected to " + websocketAddr);
    }
    websocket.onerror = function(ev) {
        console.log(ev);
        logError("websocket error: " + ev.data);
    }
    websocket.onmessage = function(ev) {
        logText("received message: " + ev.data);
    }
})()

form.addEventListener('submit', function(event){
    event.preventDefault(); // don't trigger form submit
    
    // make sure the socket is supported and ready!
    if (!websocketSupported) {
        logError("websocket is not supported by this browser");
        return;
    }
    if (!websocketReady) {
        logError("websocket is not ready");
        return;
    }

    // send a value to the socket connection
    var value = input.value;
    logText("sent message: " + value);
    websocket.send(value);
});

function logError(text) {
    var p = document.createElement('p');
    p.style.color = "red";
    p.append(document.createTextNode(text));
    results.append(p);    
}

function logText(text) {
    var p = document.createElement('p');
    p.append(document.createTextNode(text));
    results.append(p);
}

function logInfo(text) {
    var p = document.createElement('p');
    p.style.color = "green"
    p.append(document.createTextNode(text));
    results.append(p);
}


