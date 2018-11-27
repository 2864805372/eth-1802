pragma solidity ^0.4.18;

// 异常实例
contract ex {
    mapping(string=>uint) nameToBalance;
    // 插入数据
    function insert(string name, uint balance) {
        nameToBalance[name] = balance;
    }

    // 查找
    function getBalance(string name) view returns(uint) {
        if (nameToBalance[name] == 0) {
            throw; //抛出异常
        }
        if (bytes(name).length == 0) {
            throw;
        }
        return nameToBalance[name];
    }

    //
    function get(string name) {
        if getBalance(name) == 0 {
            
        }
    } 
}