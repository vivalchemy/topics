import { defineConfig } from "drizzle-kit"

export default defineConfig({
  schema: "./src/lib/drizzle/schema.ts",
  out: "./src/lib/drizzle/migrations",
  dialect: "postgresql",
  dbCredentials: {
    url: "postgres://postgres:postgres@localhost:5432/multitenant",
  },
  verbose: true,
  strict: true
}
)
