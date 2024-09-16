import { and, eq } from "drizzle-orm";
import { db } from "$lib/drizzle/db";
import { OrganizationTable, OrganizationUserTable, UserTable } from "$lib/drizzle/schema";

export async function insertUser(name: string, username: string, password: string) {
  const insertedUser = await db.insert(UserTable).values({
    name: name,
    username: username,
    password: password,
  }).returning();

  return insertedUser;
}

export async function insertOrganization(name: string) {
  const insertedOrganization = await db.insert(OrganizationTable).values({
    name: name,
  }).returning();

  return insertedOrganization;
}

export async function insertOrganizationUser(organizationId: string, userId: string) {
  const insertedOrganizationUser = await db.insert(OrganizationUserTable).values({
    organization_id: organizationId,
    user_id: userId,
  }).returning();

  return insertedOrganizationUser;
}

export async function getUserAndOrganizations(username: string, password: string) {
  // 1. Fetch user by username and password
  const user = await db.select()
    .from(UserTable)
    .where(and(eq(UserTable.username, username), eq(UserTable.password, password)))
    .limit(1)
    .then(rows => rows[0]);

  if (!user) {
    throw new Error('Invalid username or password');
  }

  // 2. Fetch organizations the user belongs to
  const organizations = await db.select({
    id: OrganizationTable.id,
    name: OrganizationTable.name,
  })
    .from(OrganizationUserTable)
    .innerJoin(OrganizationTable, eq(OrganizationUserTable.organization_id, OrganizationTable.id))
    .where(eq(OrganizationUserTable.user_id, user.id));

  // 3. Return user details and organizations
  return {
    user: {
      id: user.id,
      name: user.name,
      username: user.username,
    },
    organizations: organizations,
  };
}
