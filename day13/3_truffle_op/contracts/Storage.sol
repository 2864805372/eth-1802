pragma solidity ^0.4.18;

contract Storage {
  uint storedData;
  function set(uint data) {
    storedData = data;
  }

  function get()  constant returns (uint) {
    return storedData;
  }
}
