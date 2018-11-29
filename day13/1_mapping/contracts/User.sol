pragma solidity ^0.4.18;
import "./Iterable_mapping.sol";
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
