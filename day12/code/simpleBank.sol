pragma solidity ^0.4.18;

// 简单银行合约
contract SimpleBank{
    mapping(address=>uint) balances; // 存钱的数据库

    address public owner; // 所有人
    event LogDepositMade(address accountAddress, uint amount);
    // 构造函数
    function SimpleBank() {
        owner = msg.sender;
    }
    // 充值 payable
    function deposit() payable public returns(uint) {
        balances[msg.sender] += msg.value;
        LogDepositMade(msg.sender,msg.value);
        return balances[msg.sender];// 返回余额
    }
    // 取钱
    function withDraw(uint withDrawAmount) public returns(uint) {
        if(balances[msg.sender] >= withDrawAmount) {
            balances[msg.sender] -= withDrawAmount;
        }
        return balances[msg.sender];
    }
    // 查询余额
    function Balances() constant returns(uint) {
        return balances[msg.sender];
    }

    // fallback
    function() payable {
        throw;
    }
}