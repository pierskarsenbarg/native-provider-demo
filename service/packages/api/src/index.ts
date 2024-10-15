import { logger } from "hono/logger";
import { OpenAPIHono } from '@hono/zod-openapi';
import { bearerAuth } from 'hono/bearer-auth'
import { swaggerUI } from '@hono/swagger-ui'
import { OrganisationApi } from "./organisation";
import { TeamApi } from "./team";
import { UserApi } from "./user";

const authToken = "hunter2";

const app = new OpenAPIHono();

app
  .use(logger(), async(c, next) => {
    c.header("Cache-Control", "no-store");
    return next();
  })
  .use("/api", bearerAuth({token: authToken}));

app.openAPIRegistry.registerComponent('securitySchemes', 'Bearer', {
  type: 'http',
  scheme: 'bearer',
})

app
  .route("/organisation", OrganisationApi.route)
  .route("/team", TeamApi.route)
  .route("/user", UserApi.route);


app.doc("/doc", {
  openapi: '3.0.0',
  info: {
    version: '1.0.0',
    title: 'Native Provider API',
  },
})

app.get('/ui', swaggerUI({ url: '/doc' }))

export default { 
  port: 3000, 
  fetch: app.fetch, 
}
