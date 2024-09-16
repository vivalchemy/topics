import type { InferInsertModel, InferSelectModel } from "drizzle-orm";
import { UserTable } from "./schema";

// there are two types insert user and select user since during insert some vals can be null but during seelct they can't
export type InsertUser = InferInsertModel<typeof UserTable>
export type SelectUser = InferSelectModel<typeof UserTable>
