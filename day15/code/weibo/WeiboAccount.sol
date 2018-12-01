pragma solidity ^0.4.18;

// 个人操作平台
contract WeiboAccount {
    // 微博结构
    struct _weibo{
        uint timestamp; //微博发送时间
        string content; // 发送的微博内容
    }
    // 发送微博数量
    uint _numberOfWeibos;
    // 微博账户所有者
    address _adminAddress;
    // 通过ID进行查找
    mapping (uint => _weibo) _mapWeibos;
    // 构造函数
    function WeiboAccount(){
        
        _adminAddress = msg.sender;
        _numberOfWeibos = 0;
    }

    // 修饰符
    modifier onlyAdminAccount {
        require(_adminAddress == msg.sender);
        _;
    }

    // 发送微博
    function sendWeibo(string _content) {
        // 长度限制,微博内容不能为空，并且长度不能超过160
        require(bytes(_content).length <= 160 && bytes(_content).length != 0);
        // 微博发送时间
        _mapWeibos[_numberOfWeibos].timestamp = now;
        // 微博发送内容
        _mapWeibos[_numberOfWeibos].content = _content;
        // 更新已发送微博总数
        _numberOfWeibos++;
    }
    // 根据ID查找微博内容
    function getContentOfId(uint id) constant returns(uint timestamp,string content) {
        timestamp = _mapWeibos[id].timestamp;
        content = _mapWeibos[id].content;
    }
    // 返回最新的一条微博
    function getLatestWeibo() constant returns(uint timestamp, string content) {
        timestamp = _mapWeibos[_numberOfWeibos-1].timestamp;
        content = _mapWeibos[_numberOfWeibos-1].content;
    }
    // 返回发送的微博总数
    function getNumberOfWeibos() constant returns(uint numberOfWeibos) {
        numberOfWeibos = _numberOfWeibos;
    }

    // 返回微博账户的所有者
    function getAdminOfWeibo() constant returns(address) {
        return _adminAddress;
    }
    
    // 摧毁合约
    function adminDeleteAccount() onlyAdminAccount {
        selfdestruct(_adminAddress);
    }
}
