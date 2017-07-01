const Router = require("koa-router");
const { getUser, getAllUsers } = require("./handlers");

/**
 * User service router middleware
 * @param  {Koa} app the koa app to define routes for
 * @return {null}    does not return anything, it modifies the app.
*/
module.exports = (app) => {
  const router = new Router();

  router.get("/:id/", getUser);
  router.get("/", getAllUsers);
  router.get(/.*/, ctx => { ctx.throw(404, "Not found")});

  app.use(router.routes())
     .use(router.allowedMethods());
};
