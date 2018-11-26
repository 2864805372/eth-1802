pragma solidity ^0.4.18;

// 函数的可见性/权限
contract Chmod {
    // 私有函数
    function privateFunc() view private returns(string) {
        return "private func";
    }

    // 在当前合约内调用私有函数
    function callPrivateFunc() view public returns(string) {
        return privateFunc();
    }

    // 公有函数
    function publicFunc() view public returns(string) {
        return "public func";
    }
    // 在当前合约内调用公有函数
    function callPublicFunc() view public returns(string) {
        return publicFunc();
    }
    // 内部函数
    function internalFunc() view internal returns(string) {
        return "internal func";
    }
    // 在当前合约进行内部函数的调用
    function callInternalFunc() view public returns(string) {
        return internalFunc();
    }
    // 外部函数
    function externalFunc() view external returns(string) {
        return "external func";
    }
    // 在当前合约中进行外部函数的调用
    // 报错：不能直接在合约内部调用外部函数，只能通过this
    function callExternalFunc() view public returns(string) {
        return externalFunc();
    }
}

// 外部调用,相当于是在当前合约的上下文之外进行调用
contract externalCall {
    Chmod ch = new Chmod(); // 合约实例
    // 调用内部函数
    function callInternalFunc() view public returns(string) {
        return ch.internalFunc(); // 报错，非法调用
    }
    // 调用外部函数
    function callExternalFunc() view public returns(string) {
        return ch.externalFunc();
    }
    // 调用公有函数
    function callPublicFunc() view public returns(string) {
        return ch.publicFunc();
    }
    // 调用私有函数
    function callPrivateFunc() view public returns(string){
        return ch.privateFunc(); // 报错，私有函数不能在外部进行调用
    }
}

// 透过继承调用
contract childChmod is Chmod{
    // 调用私有函数
    function callPrivateFunc() view public returns(string) {
        return privateFunc();
    }

    // 调用公有函数
    function callPublicFunc() view public returns(string) {
        return publicFunc();
    }

    // 内部函数的调用
    function callInternalFunc() view public returns(string) {
        return internalFunc();
    }
    // 外部函数的调用
    function callExternalFunc() view public returns(string) {
        return externalFunc();
    }

}

// 继承
// 父合约
contract Father {
    uint x;
    function Father(uint _x) {
        x = _x;
    }

    function returnName() view public returns(string) {
        return("Father");
    }
}

// 母合约
contract Mother {
    uint public x;
    function Mother(uint _x) {
        x = _x;
    }
    function returnName(string str) view public returns(string) {
        return "mother";
    }
}

// 子合约
contract child is Father(10), Mother(20){
    function call() view public returns(string) {
        return returnName();
    }

    function getx() view public returns(uint) {
        return x;
    }
}

// 抽象
contract AbstractContract {
    // 抽象函数
    function someAbstractFunction(uint x); 
}

// 抽象的子合约1
contract add is AbstractContract {
    uint sum;
    function someAbstractFunction(uint x) {
        sum = x + x;
    }

    function getSum() view public returns(uint) {
        return sum;
    }
}

// 抽象的子合约2
contract mul is AbstractContract{
    uint mult;
    function someAbstractFunction(uint x) {
        mult =  x * x;
    }

    function getMul() view public returns(uint) {
        return mult;
    }
}

// 库
// 用于运算的库
library Alg {
    // add
    function add(uint x, uint y) public returns(uint) {
        return x + y;
    }

    // sub
    function sub(uint x, uint y) public returns(uint) {
        return x - y;
    }

    // mul 
    function mul(uint x, uint y) public returns(uint) {
        return x * y;
    }

    // div
    function div(uint x, uint y) public returns(uint) {
        return x / y;
    }
}

// 调用库
contract CallLibrary {
    function add(uint x, uint y) view public returns(uint) {
        return Alg.add(x, y);
    }

    function sub(uint x, uint y) view public returns(uint) {
        return Alg.sub(x, y);
    }
}


// 实现一个数组的查找库，如果存在指定元素，返回该元素索引，否则返回-1
library Search {
    function indexOf(int[] storage self, int value) returns(int) {
        for(int i = 0; i < int(self.length); i++) {
            if (self[uint(i)] == value) {
                return i;
            }
        }
        return -1;
    }
}

// using for 
contract Arr {
    using Search for int[];
    int[] data;
    // 追加元素
    function append(uint value) {
        data.push(value);
    }

    // 查找
    function contains(int value) view public returns(bool) {
        if(-1 == data.indexOf(value)) {
            return false;
        } else {
            return true;
        }
    }
}

// using for map
library Set {
    struct Data {
        mapping(uint => bool) flags;
    }

    // 插入
    function insert(Data storage self, uint value) returns(bool){
        if (self.flags[value]) {
            return false;
        } else {
            self.flags[value] = true;
            return true;
        }
    }

    // 删除
    function remove(Data storage self, uint value) returns(bool){
        if(! self.flags[value]) {
            return true;
        } else {
            self.flags[value] = false;
            return self.flags[value];
        }
    }
    // 包含
    function contains(Data storage self, uint value) returns(bool) {
         return self.flags[value];
    }
}

contract useSet {
    using Set for Set.Data;
    Set.Data data;// 结构体
    // 插入数据
    // 只传递了一个参数，是因为对于using for 来说类型实例对象会被默认当成函数的第一个参数
    function insert(uint value) returns(bool) {
        return data.insert(value);
    }

    // 包含
    function contain(uint value) returns(bool) {
        return data.contains(value);
    }
}