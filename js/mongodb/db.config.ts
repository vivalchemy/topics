//connecting to mongodb without mongoose. Without mongoose you essentially cannot create schemas as mongodb itself is schemaless
//import { MongoClient, ServerApiVersion } from "mongodb";
//
//const uri = process.env.MONGODB_URI || "";
//const client = new MongoClient(uri, {
//  serverApi: {
//    version: ServerApiVersion.v1,
//    strict: true,
//    deprecationErrors: true,
//  },
//});
//
//try {
//  // Connect the client to the server
//  await client.connect();
//  // Send a ping to confirm a successful connection
//  await client.db("admin").command({ ping: 1 });
//  console.log(
//    "Pinged your deployment. You successfully connected to MongoDB!"
//  );
//} catch (err) {
//  console.error(err);
//}
//let db = client.db("sample_mflix");
//
//export default db;
//

import mongoose from 'mongoose';

const uri = process.env.MONGODB_URI || "";

mongoose.connect(uri, {
  useNewUrlParser: true,
  useUnifiedTopology: true,
}).then(() => {
  console.log('Connected to MongoDB');
}).catch((err) => {
  console.error('Error connecting to MongoDB:', err);
});

export default mongoose;
