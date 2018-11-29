// truffle 测试文件

var Storage = artifacts.require("Storage");

contract('Storage', function(acc){
  it('get storedData', function(){
    return Storage.deployed().then(function(instance) {
      return instance.get.call();
    }).then(function(storedData) {
      assert.equal(storedData,0,"storedData");
    });
  });
});
