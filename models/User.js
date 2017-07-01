const Model = require("objection").Model;

module.exports = class User extends Model {
  static get tableName() {
    return "users";
  }

  $afterGet() {
    this.$omit("password");
  }

  static get jsonSchema() {
    return {
      type: "object",
      required: ["name", "email"],

      properties: {
        id: { type: "string"},
        name: { type: "string" },
        email: { type: "string" },
        password: { type: "string" }
      }
    };
  }
};
