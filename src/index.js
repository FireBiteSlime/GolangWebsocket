var output = document.getElementById("output");
var socket = new WebSocket("ws://localhost:8080/todo");

socket.onopen = function () {
  output.innerHTML += "Status: Connected\n";
};

socket.onclose = function () {
    output.innerHTML += "Status: Closed\n";
  };

socket.onmessage = function (e) {
  output.innerHTML += "\nServer: " + e.data + "\n";
};

function send(processedInput) {
  socket.send(processedInput);
  
}

function taskOne() {
    var inputForTaskOne = document.getElementById("inputTaskOne");
    send('1' + inputForTaskOne.value + ' ');
    inputForTaskOne.value = '';
}

function taskTwo() {
    var inputForTaskTwoF = document.getElementById("inputTaskTwoF");
    var inputForTaskTwoS = document.getElementById("inputTaskTwoS");
    send('2' + inputForTaskTwoF.value + inputForTaskTwoS.value);
    inputForTaskTwoF.value = '';
    inputForTaskTwoS.value = '';
}

function taskThree() {
    var inputForTaskThree = document.getElementById("inputTaskThree");
    send('3' + inputForTaskThree.value);
    inputForTaskThree.value = '';
}