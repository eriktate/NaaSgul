var testSocket = new WebSocket("ws://api.naasgul.io/ws");

testSocket.onopen = function(event) {
    testSocket.send("THIS IS A TEST!");
}

testSocket.onmessage = function(event) {
    console.log(event.data);
}
