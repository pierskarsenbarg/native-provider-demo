import { OpenAPIHono, createRoute, z } from "@hono/zod-openapi";
import { User } from "core/user";
import { Result } from "./common";

export module UserApi {
  export const UserSchema = z
    .object(User.Info.shape)
    .openapi("User");

  export const UserId = z.object({
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
                schema: Result(UserSchema.array()),
              },
            },
            description: "Returns a list of users",
          },
        },
      }),
      async (c) => {
        return c.json(
          {
            result: await User.list(),
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
          params: UserId,
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
                schema: Result(UserSchema),
              },
            },
            description: "Returns order",
          },
        },
      }),
      async (c) => {
        const userIdParam = c.req.param("id");

        const schemaToTest = z.object({
          value: z.coerce.number(),
        });

        const validation = await schemaToTest.safeParseAsync({
          value: userIdParam,
        });

        if (!validation.success) {
          return c.json(
            {
              error: "Invalid user id",
            },
            400
          );
        }

        const userId = z.coerce.number().parse(userIdParam);
        const user = await User.getById(userId);
        if (user === undefined) { return c.json({error: "User not found"}, 404); }
        return c.json(
          {
            result: user,
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
                schema: z.object({ name: z.string(), teamId: z.number() }),
              },
            },
          },
        },
        responses: {
          201: {
            description: "",
            content: {
              "application/json": {
                schema: Result(UserSchema),
              },
            },
          },
        },
      }),
      async (c) => {
        const user = c.req.valid("json");
        return c.json(
          {
            result: await User.create(user.name, user.teamId),
          },
          201
        );
      }
    );
}
