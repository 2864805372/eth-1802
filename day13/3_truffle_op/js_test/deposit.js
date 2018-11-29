// 合约充值
var Web3 = require("web3"); // 获取web3对象

var contract = require("truffle-contract")// 获取合约抽象实例

// 获取交互数据
var data = require("../build/contracts/Deposit.json");

// 新建实例对象
var Deposit = contract(data);

// 新建一个provider对象
var provider = new Web3.providers.HttpProvider("http://127.0.0.1:8545");

Deposit.setProvider(provider);
var depositInstance ;
Deposit.deployed().then(function(instance){
  depositInstance = instance;
  return depositInstance.getBalance.call();
}).then(result=> {
  depositInstance.sendTransaction({from:Deposit.web3.eth.accounts[0],
    value:Deposit.web3.toWei(1.5, 'ether')}).then(result=> {
      return depositInstance.getBalance.call();
    }).then(result=>{
      console.info(result.toString());
    }).catch(err => {
      console.info(err);
    })
})
