import { pgEnum, pgTable, primaryKey, uuid, varchar } from "drizzle-orm/pg-core";

// .md
//### Organization
//- _id
//- name
//- owner_id -> [[#User]]._id
//### User
//- _id
//- name
//- username
//- password
//### Organization <-> User(pk is combined)
//- organization_id
//- user_id
//- role if the role is required



// use if role is required
//export const UserRole = pgEnum("role", ["owner", "co-owner", "manager", "guest"])

export const OrganizationTable = pgTable("organizations", {
  id: uuid("id").defaultRandom().primaryKey(),
  name: varchar("name", { length: 255 }).notNull(),
})

export const UserTable = pgTable("users", {
  id: uuid("id").defaultRandom().primaryKey(),
  name: varchar("name", { length: 255 }).notNull(),
  username: varchar("username", { length: 255 }).notNull(),
  password: varchar("password", { length: 255 }).notNull(),
})

export const OrganizationUserTable = pgTable("organization_user", {
  organization_id: uuid("organization_id").references(() => OrganizationTable.id),
  user_id: uuid("user_id").references(() => UserTable.id),
  //role: UserRole("role").default("guest"),
}, table => {
  return {
    pk: primaryKey({ columns: [table.organization_id, table.user_id] }),
  }
})
