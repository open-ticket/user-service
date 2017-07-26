const Router = require("koa-router");
const {
  getAllUsers,
  createUser,
  patchUser,
  deleteUser,
  updatePassword,
  validatePassword
} = require("./handlers");

/**
 * User service router middleware
 * @param  {Koa} app the koa app to define routes for
 * @return {null}    does not return anything, it modifies the app.
*/
module.exports = (app) => {
  const router = new Router();

  router.get("/", getAllUsers);
  router.post("/", createUser);
  router.patch("/", patchUser);
  router.delete("/", deleteUser);
  router.post("/validatePassword", validatePassword);
  router.post("/updatePassword", updatePassword);
  router.get(/.*/, ctx => { ctx.throw(404, "Not found")});

  app.use(router.routes())
     .use(router.allowedMethods());
};
