import { OpenAPIHono, createRoute, z } from "@hono/zod-openapi";
import {Organisation} from "core/organisation";
import { Result } from "./common";
import type { Context } from "hono";

export module OrganisationApi {
    export const OrganisationSchema = z
        .object(Organisation.Info.shape)
        .openapi("Organisation");

    export const route = new OpenAPIHono().openapi(
        createRoute({
            security: [{ Bearer: [] }],
            method: "get",
            path: "/",
            responses: {
                200: {
                    content: {
                        "application/json": {
                            schema: Result(OrganisationSchema.array())
                        }
                    },
                    description: "Returns a list of organisations",
                }
            }
        }),
        async (c: Context) => {
            return c.json({
                result: await Organisation.list()
            }, 200)
        }
    )
    .openapi(
        createRoute({
            security: [{
                Bearer: []
            }],
            method: "post",
            path: "/",
            request: {
                body: {
                    content: {
                        "application/json": {
                            schema: Organisation.Info
                        }
                    }
                }
            },
            responses: {
                201: {
                    description: "",
                    content: {

                    }
                }
            }
        })
    )
    
    ;
}