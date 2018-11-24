pragma solidity >=0.4.22 <0.6.0;

// 数据位置转换
// contract dataLocal {
//     uint x = 5;
//     uint[5] s = [1,2,3,4,5];
//     // 1. memory -> memory
//     // function gmemory(uint a) view public returns(uint) {
//     //     a = m();
//     //     a = 2;
//     // }

//     // function m() view public returns(uint) {
//     //     return x;
//     // }

//     // memory -> storage
//     function mToS() view public returns(uint, uint) {
//         uint memory b = 10;
//         x = b;
//         return (x, b);
//     }
// }


// 枚举
contract ENUMS {
    // 自定义类型
    enum status {Created, Unlocked, Locked}
    function getStatus() view returns(status) {
        status defaultStatus = status.Locked;
        return defaultStatus;
    }
}

// 结构体
contract Struct {
     // person
    struct Person {
        string name;
        int age;
        string sex;
    }
    // bank
    // 嵌套
    struct Bank {
        address addr;
        uint balance;
        Person p;
    }
   
    // 初始化person
    Person p1 = Person({
        name:"troytan",
        age:18,
        sex:"male"
    });
    // 初始化
    Bank b1 = Bank({
        addr:0x0,
        balance:100,
        p:p1
    });

    function getBank() view public returns(address, uint, string) {
        return (b1.addr, b1.balance,b1.p1.name);
    }
    
    // 根据参数的定义顺序,不需要再添加"{}"
    Person p2 = Person({"troytan",20,"male"});
    function getPerson()returns(string, int, string) {
        return(p2.name, p2.age, p2.sex);
    }

}

// 定长数组
contract StaticArray {
    // 定长数组的定义以及初始化方式
    uint[5] a = [1,2,3,4,5];
    // 数据相关操作
    function getA() view returns(uint[5]) {
        return a;
    }
    // 获取数组长度
    function getLength() view returns(uint) {
        return a.length;
    }
    // 查找指定下标的元素
    function query(uint index) view returns(uint) {
        return a[index];
    }
    // 修改数组指定下标的元素值
    function update(uint index, uint value) {
        a[index] = value;
    }
}
// 不定长数组
contract DAarr {
    // 初始化方式一
    uint [] a = [1,2,3,4,5];
    // 初始化方式二
    // uint []a = new uint[](10);
    function getA() view returns(uint[]) {
        return a;
    }
    // 追加元素
    function append(uint value) {
        a.push(value);
    }
    // 修改长度
    function changeLength(uint len) {
        a.length = len;
    }
    // 获取长度
    function getLength() view returns(uint) {
        return a.length;
    }
}

// 数组字面量
contract constAarray {
    
    function g() view returns(uint[3]) {
        // 字面量表达式
        // 注意如果传递的参数取值与函数中默认的取值范围不同，需要转换
        return f([uint256(1),2,3]);
    }

    function f(uint[3] array) return(uint[3]) {
            return array;
    }
}

// 二维数组
contract TArray {
    // 二维数组的定义
    // 4行3列
    uint [3][4] a;

    // 修改值
    function update(uint c, uint r, uint value) {
        uint[c][r] = value;
    }
    // 返回二维数组
    function getTArray() view returns(uint[3][4]) {
        return a;
    }
    // x : 行数为5，列数不定长
    uint [][5]x;
    // uint [5][]y;
    function pushData(uint value) {
        x[0].push(value);
    }

    function getx() view returns(uint[][5]) {
        return x;
    }
}

// 字典
contract Mapping {
    // 声明方式
    mapping(address=>uint) balances;
    // insert
    function insert(address addr, uint amount) {
        balances[addr] = amount;
    }
    // query
    function get(address addr) view public returns(uint) {
        return balances[addr];
    }

    // delete
    function deleteMapping(address addr) {
        delete balances[addr];
    }

    // mapping嵌套
    mapping(address=>mapping(address=>uint)) curs;
    function insetM(address addr, uint amount) {
        curs[addr][addr] = amount;
    }

    function getMapping(address addr) view returns(uint) {
        return curs[addr][addr];
    }
}
// 隐式类型转换
contract YTrans {
    uint8 a = 1;
    uint16 b = 2;
    // 低位转高位
    function atob() view public returns(uint16) {
        b = a;
        return b;
    }

    // 高位转低位(不允许)
    function btoa() view public returns(uint8) {
        a = b;
        return a;
    }

    // uint160->address
    uint160 d = 100000000;
    address to;
    function utoa() view public returns(address) {
        to = d;
        return to;
    }
}

// 显示类型转换
contract XTrans {
    uint a = 256;
    uint8 c = 1;
    int b = -1;
    // uint -> int
    function utoi() view public returns(int) {
        b = a;
        return b;
    }

    // int->uint
    function itou() view public returns(uint) {
        a = uint(b);
        return a;
    }
    // 高位转低位
    function htol() view public returns(uint) {
        c = uint8(a);
        return c;
    }
}