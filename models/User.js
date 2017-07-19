/* eslint "class-methods-use-this": 0, "no-param-reassign": 0 */

const Model = require("objection").Model;
const bcrypt = require("bcrypt");

class User extends Model {

  static get tableName() {
    return "users";
  }

  async checkPassword(password) {
    return await bcrypt.compare(password, this.password);
  }

  $afterGet() {
    this.$omit("isDeleted");
  }

  $afterInsert() {
    this.$afterGet();
  }

  $beforeInsert () {
    // throw error if passwords don't match.
    if (this.password) {
      return bcrypt.hash(this.password, 10)
      .then(hash => {
        this.password = hash;
        this.passwordConfirm = undefined;
      });
    }
  }

  $beforeValidate(jsonSchema, json) {
    if (json.password !== json.passwordConfirm) {
      throw new Model.ValidationError(`Passwords don't match!`);
    }
  }

  static get jsonSchema() {
    return {
      type: "object",
      required: ["name", "email"],

      properties: {
        id: { type: "string"},
        name: { type: "string" },
        email: { type: "string" },
        password: { type: "string" },
        isDeleted: { type: "boolean", default: false }
      }
    };
  }
};


module.exports = User;
