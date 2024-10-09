import { sqliteTable, int, text } from "drizzle-orm/sqlite-core";

export const organisationTable = sqliteTable("organisation", {
    id: int("org_id").primaryKey({autoIncrement: true}),
    name: text("name").notNull()
});





