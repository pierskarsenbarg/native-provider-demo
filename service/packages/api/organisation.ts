import { OpenAPIHono, createRoute, z } from "@hono/zod-openapi";
import { Organisation } from "core/organisation";
import { Result } from "./common";

export module OrganisationApi {
  export const OrganisationSchema = z
    .object(Organisation.Info.shape)
    .openapi("Organisation");

  export const OrganisationId = z.object({
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
                schema: Result(OrganisationSchema.array()),
              },
            },
            description: "Returns a list of organisations",
          },
        },
      }),
      async (c) => {
        return c.json(
          {
            result: await Organisation.list(),
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
          params: OrganisationId,
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
                schema: Result(OrganisationSchema),
              },
            },
            description: "Returns order",
          },
        },
      }),
      async (c) => {
        const orgIdParam = c.req.param("id");

        const schemaToTest = z.object({
          value: z.coerce.number(),
        });

        const validation = await schemaToTest.safeParseAsync({
          value: orgIdParam,
        });

        if (!validation.success) {
          return c.json(
            {
              error: "Invalid organisation id",
            },
            400
          );
        }

        const orgId = z.coerce.number().parse(orgIdParam);
        const org = await Organisation.getById(orgId);
        if (org === undefined) { return c.json({error: "Organisation not found"}, 404); }
        return c.json(
          {
            result: org,
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
                schema: z.object({ name: z.string() }),
              },
            },
          },
        },
        responses: {
          201: {
            description: "",
            content: {
              "application/json": {
                schema: Result(OrganisationSchema),
              },
            },
          },
        },
      }),
      async (c) => {
        const name = c.req.valid("json");
        return c.json(
          {
            result: await Organisation.create(name.name),
          },
          201
        );
      }
    );
}
