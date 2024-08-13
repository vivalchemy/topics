import express from "express";
import http from "http";
import { Server } from "socket.io";

const PORT = 9000;

const app = express(); // this will handle all the https?:// html requests
const server = http.createServer(app);
const io = new Server(server); // this will handle all the ws:// websocket requests

app.use(express.static("./public"));

// Socket.io
io.on("connection", (socket) => {
  console.log("A new user has connected", socket.id);
  socket.on("client-message", (message, room) => {
    console.log("Message from client: ", socket.id, "Message: ", message);
    //io.emit("server-message", message); // to send to all clients including the sender
    //socket.broadcast.emit("server-message", message); // send to all except self clients

    // this will only send to the room you are connected to
    // rooms are basically the ids of the people you are connected to
    // can be one can be many. mostly they are one
    // the below if-else is for single person only
    if (room || room !== "") {
      socket.to(room).emit("server-message", socket.id, message); // this will by default asume that the sender don't need the message(socket.broadcast type)
    } else {
      socket.broadcast.emit("server-message", socket.id, message);
    }
  });

  // this is to join multiple people in the same room
  socket.on("join-room", (roomId, callbackFnc) => {
    console.log(roomId, socket.id);
    socket.join(roomId); // this will just keep joining you to the new rooms you won't be leaving the old rooms
    // use socket.leave(roomId) to leave the rooms
    callbackFnc("Joined room " + roomId);
  });
});

app.get("/", (req, res) => {
  return res.sendFile("/static/index.html");
});

server.listen(PORT, () => {
  console.log(`Server listening on PORT http://localhost:${PORT}`);
});
