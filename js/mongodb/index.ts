require("dotenv").config() // imports dotenv and exposes them as an object
//import express from "express"
//import cors from "cors"
import { getMovies } from "./models/movies"
// console.log(process.env)
//
//
console.log(await getMovies())
//const PORT = process.env.PORT || 3000
//const app = express()
//
//app.use(cors())
//app.use(express.json())
//
//app.get("/", (req: express.Request, res: express.Response) => {
//  console.log(req.body)
//  res.send("Hello World!")
//})
//
//app.listen(PORT, () => {
//  console.log("Server is listening on PORT :", PORT)
//})

//import { MongoClient, ServerApiVersion } from "mongodb";
//const uri = `mongodb+srv://${process.env.MONGODB_USER}:${process.env.MONGODB_PASSWORD}@test.57cbgah.mongodb.net/?retryWrites=true&w=majority&appName=test`;
//
//// Create a MongoClient with a MongoClientOptions object to set the Stable API version
//const client = new MongoClient(uri, {
//  serverApi: {
//    version: ServerApiVersion.v1,
//    strict: true,
//    deprecationErrors: true,
//  }
//});
//
//async function run() {
//  try {
//    // Connect the client to the server	(optional starting in v4.7)
//    await client.connect();
//    // Send a ping to confirm a successful connection
//    await client.db("admin").command({ ping: 1 });
//    console.log("Pinged your deployment. You successfully connected to MongoDB!");
//  } finally {
//    // Ensures that the client will close when you finish/error
//    await client.close();
//  }
//}
//run().catch(console.dir);
//
