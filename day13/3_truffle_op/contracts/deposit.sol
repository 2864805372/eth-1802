// 合约充值

pragma solidity ^0.4.18;

contract Deposit {
  // 充值事件
  event LogDeposit(address from, uint value);
  function() payable {
    LogDeposit(msg.sender, msg.value);
  }

  // 获取余额
  function getBalance() returns (uint) {
    return this.balance;
  }
}
