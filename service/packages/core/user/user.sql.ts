import { sqliteTable, int, text } from "drizzle-orm/sqlite-core";

export const userTable = sqliteTable("user", {
    id: int("user_id").primaryKey({autoIncrement: true}),
    name: text("name").notNull().unique(),
    teamId: int("team_id").notNull()
})