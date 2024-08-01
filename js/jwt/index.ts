import express from "express";
import jwt from "jsonwebtoken";
import posts from "./posts.js";
require("dotenv").config();

const app = express();

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get("/", (req, res) => {
  res.send("Hello World!");
});

app.get("/posts", authenticateToken, (req, res) => {
  console.log("posts ", posts)
  console.log(req.user)
  res.json(posts.filter((post) => post.username === req.user.username));
});

function authenticateToken(req, res, next) {
  // get the auth header
  const authHeader = req.headers['authorization']
  console.log("Headers ", req.headers)
  console.log("authHeader ", authHeader)
  // get the actual token by splitting the initial BEARER part
  const token = authHeader && authHeader.split(' ')[1]
  console.log("token ", token)
  // check if there was any token if not give error
  if (token == null) return res.sendStatus(401)
  jwt.verify(token, process.env.ACCESS_TOKEN_SECRET, (err, user) => {
    if (err) return res.sendStatus(403)
    console.log("user ", user)
    req.user = user
    next()
  })
}

app.listen(3000, () => {
  console.log("Server started on port 3000");
});

