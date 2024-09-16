import { boolean, integer, pgEnum, pgTable, primaryKey, real, serial, timestamp, unique, uniqueIndex, uuid, varchar } from "drizzle-orm/pg-core"

export const UserRole = pgEnum("role", ["admin", "student", "teacher"])

export const UserTable = pgTable("user", {
  id: uuid("id").defaultRandom().primaryKey(),
  name: varchar("name", { length: 255 }).notNull(),
  age: integer("age").notNull(),
  email: varchar("email", { length: 255 }).notNull()/*.unique() since the index applied is unique*/,
  role: UserRole("userRole").default("student").notNull()
}, table => {
  // constraint and stuff
  return {
    emailIndex: uniqueIndex("emailIndex").on(table.email),
    uniqueNameAndAge: unique("uniqueNameAndAge").on(table.name, table.age),
  }
})

export const UserPreferencesTable = pgTable("userPreferences", {
  id: uuid("id").defaultRandom().primaryKey(),
  emailUpdates: boolean("emailUpdates").notNull().default(false),
  userId: uuid("userId").references(() => UserTable.id).notNull()
})

export const PostTable = pgTable("post", {
  id: uuid("id").defaultRandom().primaryKey(),
  title: varchar("title", { length: 255 },).notNull(),
  averageRating: real("averageRating").notNull().default(0),
  createdAt: timestamp("createdAt").notNull().defaultNow(),
  updatedAt: timestamp("updatedAt").notNull().defaultNow(),
  authorId: uuid("authorId").references(() => UserTable.id).notNull()
})

export const CategoryTalble = pgTable("category", {
  id: uuid("id").defaultRandom().primaryKey(),
  name: varchar("name", { length: 255 }).notNull()
})

export const PostCategoryTable = pgTable("postCategory", {
  postId: uuid("postId").references(() => PostTable.id).notNull(),
  categoryId: uuid("categoryId").references(() => CategoryTalble.id).notNull()
}, table => {
  return {
    pk: primaryKey({ columns: [table.postId, table.categoryId] })
  }
})

export const dummy_table = pgTable("dummy_table", {
  id: serial("id").primaryKey(),
  name: varchar("name", { length: 255 }).notNull(),
  value: integer("value").notNull()
}
)
