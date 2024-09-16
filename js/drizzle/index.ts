import { sql } from "drizzle-orm";
import { db } from "./drizzle/db";
import { UserTable } from "./drizzle/schema";
import type { InsertUser, SelectUser } from "./drizzle/types";

function insertSingleUser(user: InsertUser) {
  return db.insert(UserTable)
    .values(
      user
    ).returning({
      id: UserTable.id,
      name: UserTable.name
    })
}

function insertMultipleUsers(users: InsertUser[]) {
  return db.insert(UserTable)
    .values(
      users
    ).returning({
      id: UserTable.id,
      name: UserTable.name
    }).onConflictDoUpdate({
      // this will update. It won't create a new user
      target: UserTable.email,
      set: {
        name: "Conflicting name and email"
      }
    }
    )
}

function getAllUsers() {
  // id, name, age, email, role
  // default is true until edited
  return (
    //db.query.UserTable.findMany({ columns: { id: false, name: true, email: true } }) // name, email
    //db.query.UserTable.findMany({ columns: { id: false } }) // name, age, email, role
    db.query.UserTable.findMany({
      columns: {
        id: true,
        name: true,
        email: true
      },
      extras: {
        // to run raw sql syntax: sql<return type>`raw sql`.as("column name")
        lowerCaseName: sql<string>`lower(${UserTable.name})`.as("lowerCaseName")
      }
    }) // name, age, email
    //db.select().from(UserTable) // will select everything
    //db.select({
    //  id: UserTable.id,
    //  name: UserTable.name,
    //  email: UserTable.email
    //  //UserTable.role this is not possible
    //}).from(UserTable)
  )
}

async function main() {
  //await db.delete(UserTable)
  //const user = await insertSingleUser({ name: "test", age: 12, email: "test" })
  //const users = await insertMultipleUsers([{ name: "test", age: 12, email: "test" }, { name: "VIvian", age: 12, email: "another@mail.com" }])
  const user = await getAllUsers()
  console.log(user)
  //console.log(users)
}

main()
