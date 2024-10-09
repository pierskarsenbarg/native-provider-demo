import { logger } from "hono/logger";
import { OpenAPIHono } from '@hono/zod-openapi';
import { bearerAuth } from 'hono/bearer-auth'
import { prettyJSON } from "hono/pretty-json";

const authToken = "hunter2";

const app = new OpenAPIHono();

app
  .use(logger(), async(c, next) => {
    c.header("Cache-Control", "no-store");
    return next();
  })
  .use("/api", bearerAuth({token: authToken}));

  app.use('/doc/*', prettyJSON())

app.get('/api', (c) => {
  return c.text('Hello Hono!')
});

app.openAPIRegistry.registerComponent('securitySchemes', 'Bearer', {
  type: 'http',
  scheme: 'bearer',
})

app.doc("/doc", {
  openapi: '3.0.0',
  info: {
    version: '1.0.0',
    title: 'My API',
  },
})

export default { 
  port: 3000, 
  fetch: app.fetch, 
}
