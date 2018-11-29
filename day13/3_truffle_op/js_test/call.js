var Web3 = require("web3"); // 获取web3对象

var contract = require("truffle-contract")// 获取合约抽象实例

// 获取交互数据
var data = require("../build/contracts/Storage.json");

// 新建实例对象
var Storage = contract(data);

// 新建一个provider对象
var provider = new Web3.providers.HttpProvider("http://127.0.0.1:8545");

Storage.setProvider(provider);

// 实现交互
// then的用法：将上一个函数的返回当成第一个参数传递
Storage.deployed().then(function(instance) {
  // instance ： 已经部署好的一个合约实例
  return instance.get.call();
}).then(result => {
  console.info(result.toString());
}).catch(err => {
  console.info("call failed！", err)
});
