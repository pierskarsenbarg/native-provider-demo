import {z} from "zod";
import { pipe, map, first } from "remeda";
import { eq } from "drizzle-orm";
import{db} from "../db";
import { organisationTable } from "./organisation.sql";

export module Organisation {
    export const Info = z.object({
        id: z.number(),
        name: z.string()
    });

    export type Info = z.infer<typeof Info>;

    export const create = async (name: string) => {
        var rows = await db.insert(organisationTable)
                .values({name: name}).returning();
        const result = pipe(
                rows,
                map((org): Info => ({
                    id: org.id,
                    name: org.name
                })),
                first()
            );
    
            return result as Info;
}

    export const list = async () => {
        const rows = await db.select()
                            .from(organisationTable);
        const result = pipe(
            rows,
            map((org): Info => ({
                id: org.id,
                name: org.name
            }))
        );

        return result as Info[];
    }

    export const getById = async (id: number) => {
        const rows = await db.select()
                            .from(organisationTable)
                            .where(eq(organisationTable.id, id));
        const result = pipe(
            rows,
            map((org): Info => ({
                id: org.id,
                name: org.name
            })),
            first()
        )

        return result as Info;
    }

    export const remove = async (id: number) => {
        await db
                .delete(organisationTable)
                .where(eq(organisationTable.id, id));
    }

    export const update = async (org: Info) => {
        await db
                .update(organisationTable)
                .set({
                    name: org.name
                })
                .where(eq(organisationTable.id, org.id))
    }
}