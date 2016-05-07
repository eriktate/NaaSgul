var testSocket = new WebSocket("ws://localhost:1337/ws");

testSocket.onopen = function(event) {
    testSocket.send("THIS IS A TEST!");
}

testSocket.onmessage = function(event) {
    console.log(event.data);
}
