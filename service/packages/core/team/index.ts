import {z} from "zod";
import { pipe, map, first } from "remeda";
import { eq } from "drizzle-orm";
import{db} from "../db";
import { teamTable } from "./team.sql";

export module Team {
    export const Info = z.object({
        id: z.number(),
        name: z.string(),
        organisationId: z.number()
    });

    export type Info = z.infer<typeof Info>;

    export const create = async (name: string, orgId: number) => {
        var rows = await db.insert(teamTable)
                .values({name: name, orgId: orgId}).returning();
        const result = pipe(
                rows,
                map((team): Info => ({
                    id: team.id,
                    name: team.name,
                    organisationId: team.orgId
                })),
                first()
            );
    
            return result as Info;
}

    export const list = async () => {
        const rows = await db.select()
                            .from(teamTable);
        const result = pipe(
            rows,
            map((team): Info => ({
                id: team.id,
                name: team.name,
                organisationId: team.orgId
            }))
        );

        return result as Info[];
    }

    export const getById = async (id: number) => {
        const rows = await db.select()
                            .from(teamTable)
                            .where(eq(teamTable.id, id));
        const result = pipe(
            rows,
            map((team): Info => ({
                id: team.id,
                name: team.name,
                organisationId: team.orgId
            })),
            first()
        )

        return result as Info;
    }

    export const remove = async (id: number) => {
        await db
                .delete(teamTable)
                .where(eq(teamTable.id, id));
    }

    export const update = async (team: Info) => {
        await db
                .update(teamTable)
                .set({
                    name: team.name,
                    orgId: team.organisationId,
                })
                .where(eq(teamTable.id, team.id))
    }
}