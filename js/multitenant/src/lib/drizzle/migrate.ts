import { drizzle } from "drizzle-orm/postgres-js";
import { migrate } from "drizzle-orm/postgres-js/migrator";
import postgres from "postgres";

// one way
const client = postgres("postgres://postgres:postgres@localhost:5432/multitenant" as string, { max: 1 });

async function main() {
  await migrate(drizzle(client), {
    migrationsFolder: "./src/lib/drizzle/migrations",
  });

  await client.end()
}

main()

