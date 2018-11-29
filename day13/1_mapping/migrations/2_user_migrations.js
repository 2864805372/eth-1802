var Iterator_mapping = artifacts.require("IterableMapping");
var User = artifacts.require("User");
module.exports = function(deployer) {
  deployer.deploy(Iterator_mapping);
  deployer.link(Iterator_mapping, User);
  deployer.deploy(User);
};
