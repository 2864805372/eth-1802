pragma solidity ^0.4.18;
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