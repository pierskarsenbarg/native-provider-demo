import {z} from "zod";
import { pipe, map, first } from "remeda";
import { eq } from "drizzle-orm";
import{db} from "../db";
import { userTable } from "./user.sql";

export module User {
    export const Info = z.object({
        id: z.number(),
        name: z.string(),
        teamId: z.number()
    });

    export type Info = z.infer<typeof Info>;

    export const create = async (name: string, teamId: number) => {
        var rows = await db.insert(userTable)
                .values({name: name, teamId: teamId}).returning();
        const result = pipe(
                rows,
                map((user): Info => ({
                    id: user.id,
                    name: user.name,
                    teamId: user.teamId
                })),
                first()
            );
    
            return result as Info;
}

    export const list = async () => {
        const rows = await db.select()
                            .from(userTable);
        const result = pipe(
            rows,
            map((user): Info => ({
                id: user.id,
                name: user.name,
                teamId: user.teamId
            }))
        );

        return result as Info[];
    }

    export const getById = async (id: number) => {
        const rows = await db.select()
                            .from(userTable)
                            .where(eq(userTable.id, id));
        const result = pipe(
            rows,
            map((user): Info => ({
                id: user.id,
                name: user.name,
                teamId: user.teamId
            })),
            first()
        )

        return result as Info;
    }

    export const remove = async (id: number) => {
        await db
                .delete(userTable)
                .where(eq(userTable.id, id));
    }

    export const update = async (user: Info) => {
        await db
                .update(userTable)
                .set({
                    name: user.name,
                    teamId: user.teamId,
                })
                .where(eq(userTable.id, user.id))
    }
}