import { drizzle } from "drizzle-orm/postgres-js";
import { migrate } from "drizzle-orm/postgres-js/migrator";
import postgres from "postgres";

// one way
//const client = postgres(process.env.POSTGRES_URL as string, { max: 1 });

const client = postgres({
  host: "localhost",
  user: "postgres",
  password: "postgres",
  database: "postgres",
  max: 1
});

async function main() {
  await migrate(drizzle(client), {
    migrationsFolder: "./drizzle/migrations",
  });

  await client.end()
}

main()
