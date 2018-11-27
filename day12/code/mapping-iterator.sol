// 实现mapping遍历

pragma solidity ^0.4.18;

// mapping 遍历库
library IterableMapping {
    // 自己新增相关索引
    
    // 可遍历mapping结构
    struct itmap {
        // 大小 
        uint size;
        // 存储所有keys
        KeyFlag []keys;
        // data-->mapping
        // 以mapping key为索引，index为值的map结构
        mapping(uint => IndexValue) data;
    }
    // 添加一个keys的列表
    struct KeyFlag {
        // key值
        uint key;
        // 是否删除
        bool deleted;
    }

    // value结构
    struct IndexValue {
        // key值索引进行关联
        uint keyIndex;
        // value
        uint value;
    }
    
    // 插入(修改)数据
    function insert(itmap storage self, uint key, uint value) returns(bool replaced) {
        // 优先插入数据
        uint keyIdx = self.data[key].keyIndex;
        self.data[key].value = value;
        if(keyIdx > 0) {
            return true;
        } else {
            // keys列表长度加1得到最新的keyIdx
            // (因为我们是先插入，所以采用后加)
            keyIdx = self.keys.length++;
            
            self.data[key].keyIndex = keyIdx + 1;
            self.keys[keyIdx].key = key;
            // mapping长度加1
            self.size++;
            return false;
        }
    }

    // 删除数据
    function remove(itmap storage self, uint key) returns(bool) {
        // 得到key索引
        uint keyIdx = self.data[key].keyIndex;
        if(keyIdx == 0) {
            // 表示不存在要查找的key值
            return false;
        } else {
            delete self.data[key];
            // 标记为已删除
            self.keys[keyIdx - 1].deleted = true;
            // 长度减1
            self.size--;
            return true;
        }
    }
    // 获取数据
    function iterate_get(itmap storage self, uint keyIdx) returns(uint key, uint value) {
        key = self.keys[keyIdx].key;
        value = self.data[key].value;
    }
    // 包含
    function iterate_contains(itmap storage self, uint key) returns(bool) {
        return self.data[key].keyIndex > 0;
    }

    // 获取下一个索引
    function iterate_next(itmap storage self, uint _keyIndex) returns(uint) {
        _keyIndex++;
        while(_keyIndex < self.keys.length && self.keys[_keyIndex].deleted) {
            _keyIndex++;
        }
        return _keyIndex;
    }

    // 开始遍历
    function iterate_start(itmap storage self) returns(uint keyIndex) {
        // uint(-1) = 2 ** 256 -1
        iterate_next(self,uint(-1));
    }

    // 判断循环是否要退出
    function iterate_valid(itmap storage self, uint keyIndex) returns(bool) {
        return keyIndex < self.keys.length;
    }
}

//  调用合约
contract User {
    // using for
    IterableMapping.itmap data;
    
    // 插入数据
    function insert(uint key, uint value) returns(uint) {
        IterableMapping.insert(data, key, value);
        return data.size;
    }

    // 遍历求和
    function sum() returns(uint s) {
        for (var i = IterableMapping.iterate_start(data);
        IterableMapping.iterate_valid(data, i); 
        i = IterableMapping.iterate_next(data,i)) {
            var (key, value) = IterableMapping.iterate_get(data, i);
            s += value;
        }
    }

    // 删除
    function remove(uint key) {
        IterableMapping.remove(data, key);
    }
}