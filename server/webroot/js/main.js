'use strict';

var sendButton = document.getElementById('message-send');
var messageInput = document.getElementById('message-input');
var messagesContainer = document.getElementById('messages-container');

var conn = new WebSocket('ws://' + location.host + '/wsusers');

conn.onopen = function(event) {
  console.log('Connected!', event);
}
conn.onclose = function(event) {
  console.log('Closed :(', event);
}
conn.onmessage = function (event) {
  var message = JSON.parse(event.data);
  messagesContainer.innerHTML += "<p>" + (message.user || 'Anonymous') + ": " + message.body + "</p>";
}

sendButton.addEventListener('click', function(){
  var newMessage = {body: messageInput.value};
  console.log(newMessage);
  conn.send(JSON.stringify(newMessage));
  messageInput.value = "";
  messagesContainer.innerHTML += "<p>Me: "+newMessage.body+"</p>";
});
