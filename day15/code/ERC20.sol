pragma solidity ^0.4.18;


contract Token {
    uint public totalSupply;  // 总发行量
    // 查询余额
    function balanceOf(address _owner) public constant returns(uint256 balance);
    // 转账
    function transfer(address _to, uint _value) public returns(bool success);
    // 指定从目标地址_from转账
    function transferFrom(address _from, address _to, uint _value) public returns(bool success);
    // 设置允许量值
    function approve(address _spender, uint _value) public returns(bool success);
    // 返回允许量值
    function allowance(address _owner, address _spender) public constant returns(uint remaining);

    event Transfer(address _from, address _to, uint _value);
    event Approval(address _owner, address _spender, uint _value);
}


contract TokenDemo is Token {
    // 名称
    string public name;
    // 精度
    uint8 public decimals;
    // 简称
    string public symbol;
    // 余额
    mapping(address => uint) balances;
    // 量值
    mapping(address=>mapping(address => uint)) allowed;

    // 构造函数
    function TokenDemo(uint _initialAmount, string _tokenName, uint8 _decimals, string _tokenSymbol) public{
        totalSupply = _initialAmount * 10 ** uint(_decimals); // 设置初始发行量
        name = _tokenName;
        decimals = _decimals;
        symbol = _tokenSymbol;
        balances[msg.sender] = totalSupply;
    }
    // 构造函数
    // constructor(uint _initialAmount, string _tokenName, uint8 _decimals, string _tokenSymbol) {
    //     totalSupply = _initialAmount * 10 ** uint(_decimals); // 设置初始发行量
    // }

    function balanceOf(address _owner) public constant returns(uint256 balance) {
        return balances[_owner];
    }

    function transfer(address _to, uint _value) public returns(bool success){
        require(balances[msg.sender] >= _value);
        require(_to!= 0x0);
        balances[msg.sender] -= _value;
        balances[_to] += _value;
        Transfer(msg.sender, _to, _value);
        return true;
    }
    // 指定从目标地址_from转账
    function transferFrom(address _from, address _to, uint _value) public returns(bool success){
        require(balances[msg.sender] >= _value && allowed[_from][msg.sender] >= _value);
        require(_to!= 0x0);
        balances[_from] -= _value;
        balances[_to] += _value;
        allowed[_from][msg.sender] -= _value;// 消息发送者可以从账户_from中转出的数量减少_value
        Transfer(msg.sender, _to, _value);
        return true;
    }
    // 设置允许量值
    function approve(address _spender, uint _value) public returns(bool success){
        allowed[msg.sender][_spender] = _value;
        Approval(msg.sender, _spender, _value);
        return true;
    }
    // 返回允许量值
    function allowance(address _owner, address _spender) public constant returns(uint remaining){
        return allowed[_owner][_spender];
    }

}