import { drizzle } from "drizzle-orm/postgres-js";
import * as schema from "./schema";
import postgres from "postgres";

const client = postgres("postgres://postgres:postgres@localhost:5432/multitenant")
export const db = drizzle(client, { schema, logger: true })
