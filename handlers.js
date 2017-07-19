const User = require("./models/User");
const bcrypt = require("bcrypt");

/**
 * Creates a new user using the body sent.
 * @param {Koa.ctx}  ctx Koa context
 * @return {null}    return response handled by ctx.
 */
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
    const users = await User.query()
      .where({"isDeleted": false})
      .omit(["password"]);
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
    user = await User.query().findOne("id", ctx.query.id).omit(["password"]);
  } catch (error) {
    ctx.throw(500, error.message);
  }

  if (user) {
    ctx.body = user;
  } else {
    ctx.throw(404, "User not found with that id");
  }
};

/**
 * Handles requests to patch a user by their id.
 * @param  {Koa.ctx} ctx Koa context
 * @return {null}     return response handed off to ctx
 */
const patchUser = async ctx => {
  ctx.assert(ctx.query.id, 400, "No user id provided");
  let user;

  // error if password patching
  if (ctx.request.body.password) {
    ctx.throw(400, "You cannot patch a user's password. Use the password change interface.");
  }

  try {
    user = await User
      .query()
      .omit(["password"])
      .patchAndFetchById(ctx.query.id, ctx.request.body);
    ctx.body = user;
  } catch (error) {
    ctx.throw(500, error.message);
  }


};

/**
 * Handles requests to delete users from server.
 * @param {Koa.ctx} ctx Koa context
 * @return {null}     return response handed off to ctx
 */
const deleteUser = async ctx => {
  ctx.assert(ctx.query.id, 400, "No user id provided");

  // if hard delete is true, we'll actually delete from
  // server.
  const isHardDelete = ctx.query.hard || false;
  if (isHardDelete) {
    try {
      await User.query()
        .deleteById(ctx.query.id);
        ctx.body = { success: true };
    } catch (error) {
      ctx.throw(500, error.message);
    }
  } else {
    // soft delete
    try {
      await User.query()
        .patchAndFetchById(ctx.query.id, { isDeleted: true });
      ctx.body = { success: true };
    } catch (error) {
      ctx.throw(500, error.message);
    }
  }
};

  /**
  * Handles requests to validate a user's password.
  * User ID and password are *both* sent in the body.
  */
  const validatePassword = async ctx => {
    const { userId, password } = ctx.request.body;
    if (!userId || !password) {
      ctx.throw(400, `Looks like there's an input error. Check you've provided userId + password.`);
    }

    try {
      const user = await User.query().findOne("id", userId);
      if (!user) {
        ctx.throw(404, "No user found with that ID");
      }

      ctx.body = {
        isValid: await user.checkPassword(password)
      };
    } catch (err) {
      ctx.throw(500, err.message);
    }
  };

  module.exports = {
    createUser,
    getAllUsers,
    getUser,
    patchUser,
    deleteUser,
    validatePassword
  };
