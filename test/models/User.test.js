const chai = require("chai");
const { AssertionError } = require("assert");
const User = require("../../models/User");

chai.should();

describe("create new user", () => {
  it("should create a new user", () => {
    const user = new User({
      name: "Test Name",
      email: "test@example.com"
    });

    user.name.should.equal("Test Name");
  });
});
