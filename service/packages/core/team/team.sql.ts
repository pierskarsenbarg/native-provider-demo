import { sqliteTable, int, text } from "drizzle-orm/sqlite-core";

export const teamTable = sqliteTable("team", {
    id: int("team_id").primaryKey({autoIncrement: true}),
    name: text("name").notNull().unique(),
    orgId: int("org_id").notNull()
});