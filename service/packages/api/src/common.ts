import { z } from "@hono/zod-openapi";

export function Result<T extends z.ZodTypeAny>(schema: T) {
  return z.object({
    result: schema,
  });
}

export const IdSchema = z.object({
    id: z.string().openapi({
      param: {
        name: "id",
        in: "path",
      },
    }),
  });

export const ErrorSchema = z.object({
    error: z.string(),
  });