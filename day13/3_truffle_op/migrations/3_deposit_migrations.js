var Deposit = artifacts.require("Deposit");

// deposti 迁移脚本
module.exports = function(deployer) {
  deployer.deploy(Deposit);
};
