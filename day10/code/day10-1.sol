// 导入solidity编译器
pragma solidity >=0.4.22 <0.6.0;

// 注释
contract HelloWorld {
    function hello() returns(string) {
        return "hello world";
    }
}

// 整形
contract Interger {
    // 有符号整形
    int a = 100;
    // 无符号整形
    uint b = 200;

    function getInt() public returns(int, uint) {
        return(a, b);
    }

    function add(uint x, uint y) view public returns(uint z) {
        return x - y;
    }
    
    function and(uint x, uint y) view public returns(uint) {
        return x & y;
    }
    // 乘方运算
    function square(uint x, uint y) view public returns(uint) {
        return x ** y;
    }

    uint8 s = 255;
    // 整形上溢
    function maxLimit(uint8 num) view public returns(uint8) {
        return s + num;
    }
    
}

// bool
contract Bool {
    function getBool() public returns(bool) {
        return 100 < 200;
    }
}

// 地址类型
contract Address {
    // 表达形式
    address addr = 0xca35b7d915458ef540ade6068dfe2f44e8fa733c;
    function getAddress() view public returns(address) {
        return addr;
    }

    // 地址类型比较运算
    function compare(address a, address b) view public returns(bool) {
        return a > b;
    }

    // 获取指定地址余额
    function getBalance(address addr) view public returns(uint) {
        return address.balance;
    }

    // 获取当前账号地址
    function getCurrentAddress() view public returns(address) {
        return msg.sender;
    }

    // 转账
    function sendTx() view public {
        msg.sender.send(100);
    }

    // 充值
    function deposit() payable  {
        sendTx()
    }
}

// 定长字节数组
contract Bytes{
    bytes1 a = 0x0a;
    bytes2 b = 0x0aa0;
    // 关系运算符
    function bigger() view returns(bool) {
        return a > b;
    }

    // 位运算符
    function and() view return (bytes1) {
        return a & b;
    }

    // 获取长度
    function getLength() view returns(uint) {
        return a.length;
    }
}


// 字符串
contract String {
    string name = "themoonstone";
    function getName() view public returns(string) {
        return name;
    }
    // 获取字符串长度
    function getLength() view public returns(uint) {
        return name.length;
    }

    // 对字符串取值
     function getValue() view returns(bytes1) {
        return bytes(name)[0];
    }
    // 转换
     function getNameBytes() view returns(bytes) {
        return bytes(name);
    }

    // 修改指定字符
    function changeName() {
        bytes(name)[0] = 's';
    }
}

// 智能推断var 
contract Var {
    // 使用var
    function VarValue() view returns(uint) {
        var a = 10;
        return a;
    }
    // 将函数值赋给一个var 变量
    function callFunc() view returns(uint) {
        var s = VarValue;
        return s();
    }
}

// 十六进制字面量
contract HexLi{
    function test() returns(string) {
        var a = hex"001122FF";
        return a;
    }

    function hexTobyte() view returns(bytes4,bytes1, bytes1) {
        bytes4 a = hex"001122FF";
        return (a, a[0],a[3]);
    }
}

// 常量
contract Constant {
    string constant test = "abc";
}