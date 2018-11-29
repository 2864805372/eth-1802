// 合约交互之交易

var Web3 = require("web3"); // 获取web3对象

var contract = require("truffle-contract")// 获取合约抽象实例

// 获取交互数据
var data = require("../build/contracts/Storage.json");

// 新建实例对象
var Storage = contract(data);

// 新建一个provider对象
var provider = new Web3.providers.HttpProvider("http://127.0.0.1:8545");

Storage.setProvider(provider);

var storageInstance;
// 合约交易
Storage.deployed().then(function(instance){
  storageInstance = instance;
  return storageInstance.set.sendTransaction(42,{from:Storage.web3.eth.accounts[0]});
}).then(result => {
  // result是一个对象，包含以下值
  // tx--交易哈希
  // receipt--交易对象
  // logs--在交易中触发的日志
  console.info(result);
}).then(()=>{
  return storageInstance.get.call();
}).then(result=>{
  console.info(result.toString());
}).catch(err => {
  console.info("send transaction failed! ", err);
});
