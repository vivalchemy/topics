import type { InferInsertModel, InferSelectModel } from "drizzle-orm";
import { UserTable, OrganizationTable, OrganizationUserTable } from "./schema";

export type InsertUser = InferInsertModel<typeof UserTable>;
export type SelectUser = InferSelectModel<typeof UserTable>;
export type InsertOrganization = InferInsertModel<typeof OrganizationTable>;
export type SelectOrganization = InferSelectModel<typeof OrganizationTable>;
export type InsertOrganizationUser = InferInsertModel<typeof OrganizationUserTable>;
export type SelectOrganizationUser = InferSelectModel<typeof OrganizationUserTable>;
