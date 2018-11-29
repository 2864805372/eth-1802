
// 基于指定合约地址生成一个新的合约对象

var Web3 = require("web3"); // 获取web3对象

var contract = require("truffle-contract")// 获取合约抽象实例

// 获取交互数据
var data = require("../build/contracts/Storage.json");

// 新建实例对象
var Storage = contract(data);

// 新建一个provider对象
var provider = new Web3.providers.HttpProvider("http://127.0.0.1:8545");

Storage.setProvider(provider);

Storage.at("0xac76fe28f6e8552b21c24f6d6b02e39f5f72e0fe").then(function(instance) {
  return instance.get.call();
}).then(result => {
  console.info(result.toString());
}).catch(err => {
  console.info(err);
})
