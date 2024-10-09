import { defineConfig } from "drizzle-kit";

export default defineConfig({
    dialect: "sqlite",
    schema: "./**/*.sql.ts",
    out: "./packages/core/migrations",
    dbCredentials: {
        url: "file:./db.sqlite"
    },
})