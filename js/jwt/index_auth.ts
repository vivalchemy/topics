import express from "express";
import jwt from "jsonwebtoken";
require("dotenv").config();

const app = express();

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// this is created for demo purpose only. Use a db/local storage or something to store this
let refreshTokens: string[] = [];

function generateAccessToken(username: string) {
  return jwt.sign({ username }, process.env.ACCESS_TOKEN_SECRET, {
    expiresIn: "30s",
  });
}

function generateRefreshToken(username: string) {
  return jwt.sign({ username }, process.env.REFRESH_TOKEN_SECRET);
}

app.post('/login', (req: express.Request, res: express.Response) => {
  const username: string = req.body.username;
  console.log("username ", username)
  const accessToken = generateAccessToken(username)
  const refreshToken = generateRefreshToken(username)
  refreshTokens.push(refreshToken)
  console.log("accessToken: ", accessToken, "\nrefreshToken", refreshToken)
  res.json({ accessToken, refreshToken })
});

app.post('/token', (req: express.Request, res: express.Response) => {
  console.log(req.body);
  const refreshToken = req.body.token
  if (refreshToken == null) return res.sendStatus(401)
  if (!refreshTokens.includes(refreshToken)) return res.sendStatus(403)
  jwt.verify(refreshToken, process.env.REFRESH_TOKEN_SECRET, (err: Error, user: { username: string }) => {
    if (err) return res.sendStatus(403)
    const accessToken = generateAccessToken(user.username)
    res.json({ accessToken })
  })
})


app.delete("/logout", (req: express.Request, res: express.Response) => {
  refreshTokens = refreshTokens.filter(token => token !== req.body.token)
  res.sendStatus(204)
})

app.listen(4000, () => {
  console.log("Server started on port 4000");
});

