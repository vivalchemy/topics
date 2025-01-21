//let socket = new WebSocket("ws://localhost:3000/ws"); // simulate the chat app
let socket = new WebSocket("ws://localhost:3000/stream"); // simulate the server streaming to data

socket.onmessage = function(event: MessageEvent) {
  console.log("Received by client::", event.data)
}
// this should start the socket server and receive the message

//socket.send("Hello from client") // send to /ws route 
