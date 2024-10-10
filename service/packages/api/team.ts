import { OpenAPIHono, createRoute, z } from "@hono/zod-openapi";
import { Team } from "core/team";
import { Result } from "./common";

export module TeamApi {
  export const TeamSchema = z
    .object(Team.Info.shape)
    .openapi("Team");

  export const TeamId = z.object({
    id: z.string().openapi({
      param: {
        name: "id",
        in: "path",
      },
    }),
  });

  export const route = new OpenAPIHono()
    .openapi(
      createRoute({
        security: [{ Bearer: [] }],
        method: "get",
        path: "/",
        responses: {
          200: {
            content: {
              "application/json": {
                schema: Result(TeamSchema.array()),
              },
            },
            description: "Returns a list of teams",
          },
        },
      }),
      async (c) => {
        return c.json(
          {
            result: await Team.list(),
          },
          200
        );
      }
    )
    .openapi(
      createRoute({
        method: "get",
        path: "/{id}",
        request: {
          params: TeamId,
        },
        responses: {
          404: {
            content: {
              "application/json": {
                schema: z.object({
                  error: z.string(),
                }),
              },
            },
            description: "Not found",
          },
          400: {
            content: {
              "application/json": {
                schema: z.object({
                  error: z.string(),
                }),
              },
            },
            description: "Bad request",
          },
          200: {
            content: {
              "application/json": {
                schema: Result(TeamSchema),
              },
            },
            description: "Returns team",
          },
        },
      }),
      async (c) => {
        const teamIdParam = c.req.param("id");

        const schemaToTest = z.object({
          value: z.coerce.number(),
        });

        const validation = await schemaToTest.safeParseAsync({
          value: teamIdParam,
        });

        if (!validation.success) {
          return c.json(
            {
              error: "Invalid team id",
            },
            400
          );
        }

        const teamId = z.coerce.number().parse(teamIdParam);
        const team = await Team.getById(teamId);
        if (team === undefined) { return c.json({error: "Team not found"}, 404); }
        return c.json(
          {
            result: team,
          },
          200
        );
      }
    )
    .openapi(
      createRoute({
        security: [
          {
            Bearer: [],
          },
        ],
        method: "post",
        path: "/",
        request: {
          body: {
            content: {
              "application/json": {
                schema: z.object({ name: z.string(), orgId: z.number() }),
              },
            },
          },
        },
        responses: {
          201: {
            description: "",
            content: {
              "application/json": {
                schema: Result(TeamSchema),
              },
            },
          },
        },
      }),
      async (c) => {
        const team = c.req.valid("json");
        return c.json(
          {
            result: await Team.create(team.name, team.orgId),
          },
          201
        );
      }
    );
}
