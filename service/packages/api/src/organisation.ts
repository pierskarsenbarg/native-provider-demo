import { OpenAPIHono, createRoute, z } from "@hono/zod-openapi";
import { Organisation } from "@service/core/organisation";
import { Result, IdSchema, ErrorSchema } from "./common";

export module OrganisationApi {
  export const OrganisationSchema = z
    .object(Organisation.Info.shape)
    .openapi("Organisation");

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
        security: [{ Bearer: [] }],
        method: "get",
        path: "/{id}",
        request: {
          params: IdSchema,
        },
        responses: {
          404: {
            content: {
              "application/json": {
                schema: ErrorSchema,
              },
            },
            description: "Not found",
          },
          400: {
            content: {
              "application/json": {
                schema: ErrorSchema,
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
            description: "Returns organisation",
          },
        },
      }),
      async (c) => {
        const { id } = c.req.valid("param");
        const schemaToTest = z.object({
          value: z.coerce.number(),
        });

        const validation = await schemaToTest.safeParseAsync({
          value: id,
        });

        if (!validation.success) {
          console.log(`Org id: ${id}`)
          return c.json(
            {
              error: "Invalid organisation id",
            },
            400
          );
        }

        const orgId = z.coerce.number().parse(id);
        const org = await Organisation.getById(orgId);
        if (org === undefined) {
          return c.json({ error: "Organisation not found" }, 404);
        }
        return c.json(
          {
            result: org,
          },
          200
        );
      },
      // (result, c) => {
      //   if (!result.success) {
      //     console.log(result)
      //     return c.json(
      //       {
      //         error: "Invalid organisation id thing",
      //       },
      //       400
      //     );
      //   }
      // }
    )
    .openapi(
      createRoute({
        security: [
          {
            Bearer: [],
          },
        ],
        method: "delete",
        path: "/{id}",
        request: {
          params: IdSchema,
        },
        responses: {
          404: {
            content: {
              "application/json": {
                schema: ErrorSchema,
              },
            },
            description: "Not found",
          },
          400: {
            content: {
              "application/json": {
                schema: ErrorSchema,
              },
            },
            description: "Bad request",
          },
          204: {
            description: "Organisation deleted",
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
        await Organisation.remove(orgId);
        return c.body(null, 204);
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
            description: "Organisation successfully created",
            content: {
              "application/json": {
                schema: Result(OrganisationSchema),
              },
            },
          },
          400: {
            description: "Bad request",
            content: {
              "application/json": {
                schema: ErrorSchema,
              },
            },
          },
        },
      }),
      async (c) => {
        const name = c.req.valid("json");
        try {
          return c.json(
            {
              result: await Organisation.create(name.name),
            },
            201
          );
        } catch (err) {
          return c.json(
            {
              error: "Bad request",
            },
            400
          );
        }
      }
    )
    .openapi(
      createRoute({
        security: [
          {
            Bearer: [],
          },
        ],
        method: "put",
        path: "/",
        request: {
          body: {
            content: {
              "application/json": {
                schema: z.object({ id: z.number(), name: z.string() }),
              },
            },
          },
        },
        responses: {
          200: {
            description: "Organisation successfully updated",
          },
          400: {
            description: "Bad request",
            content: {
              "application/json": {
                schema: ErrorSchema,
              },
            },
          },
        },
      }),
      async (c) => {
        const orgRes = c.req.valid("json");

        const schema = z.object({
          id: z.coerce.number(),
          name: z.coerce.string(),
        });

        const validation = await schema.safeParseAsync(orgRes);

        if (!validation.success) {
          return c.json(
            {
              error: "Invalid organisation",
            },
            400
          );
        } 

        const org = schema.parse(orgRes);
        try {
          await Organisation.update(org);
          return c.body(null, 200);
        } catch (err) {
          return c.json(
            {
              error: err,
            },
            500
          );
        }
      },
      (result, c) => {
        if (!result.success) {
          console.log(result)
          return c.json(
            {
              error: "Invalid organisation id thing",
            },
            400
          );
        }
      }
    );
}
