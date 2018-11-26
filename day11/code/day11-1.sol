pragma solidity ^0.4.18;

contract For {
    // 等差数列求和
    function getSum(uint n, uint x) public returns(uint sum) {
        for (var index = 0; index < n; index+=x) {
            sum += index;
            if (index > 50) {
                break;
            }
        }
    }
}

contract While {
    function getSum(uint n, uint x) view public returns(uint sum) {
        uint i = 0;
        while(i < n) {
            sum += i;
            i += x;
        }
    }
}

// 后测试,会优先执行一次语句块do中的代码逻辑，然后再进行条件判断
contract DoWhile {
    function getSum(uint n, uint x) view public returns(uint sum) {
        uint i = 0;
        do {
            sum += i;
            i+=x;
        }while(i<n);
    }
}

// 控制语句
contract Constructor {
    function getIf(uint x) view public returns(string) {
        if (x % 2 == 0) {
            return "偶数";
        }  else {
            return "奇数";
        }
    }

    // 三目运算符
    function getT(uint x) view public returns(string) {
        return (x % 2 == 0) ? "偶数":"奇数";
    }
}

// 全局单位
contract Global {
    function getMsg() view returns(address) {
        return msg.sender;
    }

    function getThis() view returns(address,uint) {
        return (this,this.balance);
    }

    function getTx() view returns(address, uint) {
        return (tx.origin, tx.gas);
    }

    function getBlock() view returns(uint) {
        return block.number;
    }
}

// 数据位置转换
// 引用传递
contract MemoryToMemory {
    struct S {
        string a;
        uint b;
    }

    function mtom(S s) {
        S memory tmp = s;

    }

    function call() view returns(string) {
        S memory s = S("memory", 1);
        mtom(s);
        return s.a;
    }
}

// storage->storage
// 引用传递
contract StorageToStorage {
     struct S {
        string a;
        uint b;
    }
    S s;
    function stos(S storage x) internal{
        S x;
        x.a = "other storage";
    }

    function call() view returns(string) {
        stos(s);
        return s.a;
    }
}

// memory->storage
// 值传递
contract MemoryToStorage {
    struct S {
        string a;
        uint b;
    }
    S s;
    function mtos(S memory x) {
        s = x;
        s.a = "storage";
    }
    function call() view returns(string) {
        S memory tm = S("memory", 1);
        mtos(tm);
        return tm.a;
    }
}

// storage->memory
// 值传递
contract StorageToMemory {
    struct S {
        string a;
        uint b;
    }
    S s = S("storage", 1);
    function stom() internal {
        S memory tmp = s;
        tmp.a = "memory";
    }

    function call() view returns(string) {
        stom();
        return s.a;
    }
}

// 函数
contract Func {
    uint x = 100;
    uint y = 200;
    function add() int returns(uint) {
        return x + y;
    }

    // 内部调用
    function internalCallAdd() public view returns(uint) {
        return add();
    } 

     mapping(address => uint) balances;
    function insert() public {
        balances[msg.sender] = 100;
    }
    // 多返回值
    function get() public view returns(address, uint) {
        return(msg.sender, balances[msg.sender]);
    }
}

// 函数类型
contract FunctionType {
    uint x = 100;
    // view
    function modifyx() view public returns(uint) {
        x = 50;
        g();
        return x;
    }

    function g() public view {
        
    }

    // pure
    function modifyXPure() pure public returns(uint) {
        x = 20; // 修改状态变量，报错
        return x; // 读取合约状态，报错
    }

    function getBalanceOfAddress() pure returns(uint) {
        return address(this).balance;
    }
}

// 回退函数
contract FallBackFunction{
    // 新建一个事件，把调用的数据打印
    event ExistFuncCalled(bytes data, uint256 para);
    event FallbackCalled(bytes data);
    // fallback函数
    function() {
        FallbackCalled(msg.data);
    }

    // 调用一个已经存在的函数
    function existFunc(uint256 para) {
        ExistFuncCalled(msg.data, para);
    }
    // 调用一个不存在的函数
    function callNotExistFunc() {
        bytes4 funcIdentifier = bytes4(keccak256("functionNotExist"));
        this.call(funcIdentifier);
    }
}

// 自毁函数
contract SelfDes {
    string public someValue;
    address public owner;
    // 构造函数
    function SelfDes() {
        owner = msg.sender;
    }
    // 设置一个值
    function setSomeValue(string value) {
        someValue = value;
    }

    // 调用自毁函数
    function destroyContract() {
        // 自毁函数
        selfdestruct(owner);
    }
}

// 常函数
contract ConstFunc{
    function f(uint a, uint b) constant returns(uint) {
        return a * b;
    }
}

// 访问器
// getter
contract C {
    uint public c = 100;
    function accessInternal() returns(uint) {
        return c; // 内部访问方式
    }

    function accessExternal() returns(uint) {
        return this.c();
    }
}

// 调用其它合约的访问器
contract D {
    // 实例化其它合约对象
    C c = new C();
    function getData() returns(uint) {
        return c.c();
    }
}


// 内置数学函数
contract Math {
    function mod(uint x, uint y, uint k) view public returns(uint) {
        // x + y 对k 取余
        return addmod(x, y, k);
    }

    function mmod(uint x, uint y, uint k) view public returns(uint) {
        // x * y对k取余
        return mulmod(x, y, k);
    }
}

// 加密函数
contract Sha {
    function getS(string x) returns(bytes32) {
        return sha256(x);
    }
}


// 函数修改器
contract Modify {
    address owner = msg.sender;
    // 定义
    modifier onlyOwner(){
        // 前置条件
        require(msg.sender==owner);
        _;
    }

    modifier balances() {
        require(msg.sender.balance >= 100000000);
        _;
    }

    function modifyOwner(address addr) onlyOwner() balances() view public returns(address) {
        // 转移合约的所有权
        owner = addr;
        return owner;
    }
}
