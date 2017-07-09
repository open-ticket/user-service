const User = require("./models/User");

const createUser = async ctx => {
  try {
    ctx.body = await User.query().insert(ctx.request.body);
  } catch (error) {
    ctx.throw(500, error.message);
  }
};

/**
 * Handles requests to get all users from DB.
 * @param  {Koa.ctx}  ctx Koa context
 * @return {null}     return response handled by ctx
 */
const getAllUsers = async ctx => {
  if (ctx.query.id) {
    return getUser(ctx);
  }
  try {
    const users = await User.query();
    ctx.body = users;
  } catch (error) {
    ctx.throw(500, error.message);
  }
};

/**
 * Handles requests to get a user by their id.
 * @param  {Koa.ctx} ctx Koa context
 * @return {null}     return response handed off to ctx
 */
const getUser = async ctx => {
  ctx.assert(ctx.query.id, 400, "No user id provided");

  let user;
  try {
    user = await User.query().findOne("id", ctx.query.id);
  } catch (error) {
    ctx.throw(500, error.message);
  }

  if (user) {
    ctx.body = user;
  } else {
    ctx.throw(404, "User not found with that id");
  }
};

module.exports = {
  createUser,
  getAllUsers,
  getUser,
};
