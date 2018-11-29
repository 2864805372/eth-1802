//添加一个新的合约到网络中
var Web3 = require("web3"); // 获取web3对象

var contract = require("truffle-contract")// 获取合约抽象实例

// 获取交互数据
var data = require("../build/contracts/Storage.json");

// 新建实例对象
var Storage = contract(data);

// 新建一个provider对象
var provider = new Web3.providers.HttpProvider("http://127.0.0.1:8545");

Storage.setProvider(provider);

// new：生成一个新的合约
// instance：新生成的合约实例
Storage.new({from:Storage.web3.eth.accounts[0], gas:1000000}).then(function(instance) {
  console.info(instance.address);
}).catch(err => {
  console.info(err);
});
