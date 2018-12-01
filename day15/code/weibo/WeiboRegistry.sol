pragma solidity ^0.4.18;

// 微博平台
contract WeiboRegistry {
    // 平台上微博账户数量
    uint _numberOfAccounts;
    // 管理员账号
    address _adminAccount;
    // id=>address
    mapping(uint => address) _idToAccount;
    // 账号=>昵称
    mapping(address => string) _accountToName;
    // 昵称=>账号
    mapping(string => address) _nameToAccount;

    // 修改器
    modifier onlyAdmin {
        require(msg.sender == _adminAccount);
        _;
    }

    // 构造函数
    function WeiboRgistry() {
        _numberOfAccounts = 0;
        _adminAccount = msg.sender;
    }

    /*功能函数 */

    // 注册账户
    function register(string name, address addr) {
        // 限制昵称长度
        require(bytes(name).length <= 20 && bytes(name).length != 0);
        // 地址不能为空地址
        require(addr != address(0));
        // 账号没有被注册
        require(bytes(_accountToName[addr]).length == 0);
        // 昵称没有被注册
        require(_nameToAccount[name] == address(0));


        _idToAccount[_numberOfAccounts] = addr;
        _accountToName[addr] = name;
        _nameToAccount[name] = addr;

        // 账号总数加1
        _numberOfAccounts++;
    }
    // 获取账户数量
    function getNumberOfAccount() constant returns(uint numberOfAccounts) {
        numberOfAccounts = _numberOfAccounts;
    }
    // 通过ID查找地址
    function getAccOfId(uint id) constant returns(address account) {
        account = _idToAccount[id];
    }
    // 通过地址查找昵称
    function getNameOfAccount(address acc) constant returns(string name) {
        name = _accountToName[acc];
    }
    // 通过昵称查找地址
    function getAccOfName(string name) constant returns(address account) {
        account = _nameToAccount[name];
    }
    // 摧毁合约
    function adminDeleteAccount() onlyAdmin {
        selfdestruct(_adminAccount);
    }
}