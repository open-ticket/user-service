const Koa = require("koa");
const Knex = require("knex");
const objection = require("objection");
const bodyparser = require("koa-bodyparser");
const knexConfig = require("./knexfile");
const { logger, timer, jsonify } = require("./middleware");
const router = require("./router");

require("dotenv").config();

// connect to the database.
const knex = Knex(knexConfig);
objection.Model.knex(knex);

const app = new Koa();
app.use(bodyparser());

app.use(timer);
app.use(logger);
app.use(jsonify);

// load routes
router(app);

app.listen(process.env.PORT || 3000);
