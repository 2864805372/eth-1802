1. day11复习

   1. 函数
   2. 调用方式
      1. 内部调用
      2. 外部调用
   3. 定义
      1. view : 只读，不改变合约的内部状态
      2. pure : 针对合约的内部状态，既不能写，也不能读
   4. 特殊函数
      1. 回退函数
      2. 自毁函数
      3. 常函数
      4. 内置数学函数
      5. 内置加密函数
      6. 访问器
   5. 函数修改器
      1. modifier
   6. 函数的可见性与权限
      1. 分类：private,internal,external,public
   7. solidity面向对象
      1. 继承
         1. 合约关系
         2. is
         3. 支持参数传递
         4. 支持多重继承
         5. 同名函数覆盖
      2. 抽象
         1. 抽象函数
         2. 抽象合约
      3. 库(library)
         1. 库也是一种合约
         2. using A for B
      4. import
         1. 导入源文件
         2. import "filename"

2. 事件

   1. 事件是使用EVM日志所内置的工具，关键字event
   2. 为什么要有事件
      1. 在真实环境中，发送交易调用智能合约的过程
         1. 交易发送-->打包-->执行交易。在发送交易之后，不会马上返回执行结果，只会立刻返回一个交易的哈希
      2. 事件可以继承，在合约内不能直接访问

3. 异常处理

   1. throw:如果发生异常，消耗所有gas，没有异常信息，回滚所有状态
   2. require(bool condition):自行判断，如果不满足条件也会产生异常，返回未使用的gas，回滚已修改状态
   3. assert(bool condition):如果产生异常，返回未使用的gas,回滚已修改状态
   4. revert():终止执行，回滚所有状态
   5. 注意：一般情况下，尽可能的使用require(bool condition)
   6. 在智能合约开发中，异常处理应该遵循如下规则
      1. 尽早抛出异常
      2. 在函数中，针对异常的发生，组织代码顺序
         1. 检查所有的前置条件
         2. 修改合约状态
         3. 和其他合约进行交互

4. 实现mapping遍历

   1. 实现一个mapping库，在原有mapping支持增，删，改，查基础上增加遍历功能
   2. 关键点
      1. 索引
      2. 长度
      3. 结构体

5. solidity合约手动编译部署

   1. 编写合约

   2. ```javascript
      pragma solidity ^0.4.18;
      
      contract Storage {
      	uint256 storedData;
      	function set(uint data) {
      		storedData = data;
      	}
      	function get() constant returns(uint) {
      		return storedData;
      	}
      }
      ```

   3. 手动编译合约

   4. ```javascript
      echo "var storageOutput=`solc --optimize --combined-json abi,bin,interface Storage.sol`" > storage.js
      ```

   5. cat storage.js

   6. ```json
      [
          {
              "constant":false,
              "inputs":[
                  {
                      "name":"data",
                      "type":"uint256"
                  }
              ],
              "name":"set",
              "outputs":[
      
              ],
              "payable":false,
              "stateMutability":"nonpayable",
              "type":"function"
          },
          {
              "constant":true,
              "inputs":[
      
              ],
              "name":"get",
              "outputs":[
                  {
                      "name":"",
                      "type":"uint256"
                  }
              ],
              "payable":false,
              "stateMutability":"view",
              "type":"function"
          }
      ]
      ```

6. 字段解释

   1. ABI:应用二进制接口(Application Binary Interface),是从区块链外部与合约进行交互或者合约与合约进行交互的一种标准方式，可以理解为一种编码
      1. type:方法类型：总共包括function, constructor, fallback,默认是function
      2. name:方法名
      3. inputs:方法的参数列表，一个对应的数据，数组中的每个对象都是参数说明
         1. name:参数名称
         2. type:参数类型
      4. outputs:方法的输出列表(返回值值列表)
      5. constant:布尔值，如果为true，则说明该方法不会修改合约的状态变量
      6. payable:布尔值，表示方法是否能够接收ether
      7. stateMutability:定义类型，检查是否会修改状态变量，以及是否可以接收ether
   2. bin:合约编译之后的二进制内容

7. 部署合约(部署到私链上面)

   1. 创世配置文件

   2. 生成创世区块:geth --datadir db init genesis.json

   3. 启动以太坊节点 

   4. ```
      geth --datadir db --rpc --rpcaddr=0.0.0.0 --rpccorsdomain "*" --nodiscover --maxpeers 30 --networkid 15 --rpcapi "eth,net,web3,personal,admin,txpoo,debug,miner" --mine --minerthreads 1 --etherbase "0xca35b7d915458ef540ade6068dfe2f44e8fa733c" console
      ```

   5. 为了方便，再启动一个链接连上这个节点 ：geth --datadir ./db attach ipc:./db/geth.ipc

   6. 加载storage.js文件：loadScript('/home/eth/eth1802/day12/storage.js')

   7. 获取合约的ABI以及二进制代码

      1. 获取ABI：var storageContractAbi=storageOutput.contracts['Storage.sol:Storage'].abi
      2. 解析ABI得到contract:var storageContract = eth.contract(JSON.parse(storageContractAbi))
      3. 获取bin：var storageBinCode="0x"+storageOutput.contracts['Storage.sol:Storage'].bin

   8. 部署智能合约

      1. 创建账户：personal.newAccount("123456")
      2. 设置新的挖矿地址:miner.setEtherbase(eth.accounts[0])
      3. 查看挖矿地址
      4. 查看余额：可以发现在新挖矿地址设置成功之后，挖矿奖励会直接转到新的地址上面
      5. 解锁账户：personal.unlockAccount(eth.accounts[0])
      6. 向网络中发送部署合约的交易
         1. var deployTransactionObject={from:eth.accounts[0],data:storageBinCode,gas:100000}
         2. var storageInstance = storageContract.new(deployTransactionObject)
         3. 部署：var storageInstance = storageContract.new(deployTransactionObject)
         4. 根据部署合约的交易hash查看详情:eth.getTransactionReceipt(storageInstance.transactionHash)
         5. 获取合约地址：eth.getTransaction(storageInstance.transactionHash).contractAddress

   9. 合约交互

   10. solidity总结

      1. 简单银行实现
         1. 三个基本功能
            1. 充值
            2. 取钱
            3. 查询余额
      2. 投票合约
         1. 基本功能
            1. 投票
            2. 委托

   11. web3.js

       1. 概念：与以太坊合约进行交互的javascript api
       2. 作用：以太坊节点只能识别JSON-RPC的语言，但是JSON语句写起来非常麻烦，而极其容易出错，而web3.js把JSON-PRC封装起来，让开发者只需要和JS控制台交互，不需要直面JSON语句
       3. 通常问下，在solidty的合约开发中，web3.js更多与truffle开发框架结合

   12. truffle开发框架

       1. 主要功能

          1. 内置了合约的编译，部署，二进制字节管理
          2. 脚本化，可以快速部署和迁移
          3. 网络管理，可以部署到任意数量的公共网络和私有网络
          4. 使用npm进行管理
          5. 交互式控制台用于和合约进行通信

       2. truffle具体使用

          1. 安装：npm install truffle

          2. 目录初始化:truffle init

             1. contracts:存储编写合约
             2. migrations:存放迁移部署的脚本
             3. test:存储合约测试脚本
             4. truffle.js:truffle配置文件
             5. truffle-config.js:truffle配置文件

          3. 创建合约

          4. 编译合约：truffle compile,注意，如果需要编译已经编译过的合约，使用truffle --all

          5. 部署迁移

             1. testrpc

             2. ganache:testrpc的升级版，下载地址:https://truffleframework.com/ganache

             3. 修改truffle.js，添加对应的配置

             4. ```javascript
                 module.exports = {
                   networks: {
                     development: {
                       host: "127.0.0.1",
                       port: 8545,
                       network_id: "*" // Match any network id
                     }
                   }
                 };
                ```

             5. 添加部署脚本文件：必须以数字开头

             6. ```javascript
                var Storage = artifacts.require("./Storage.sol");
                
                module.exports = function(deployer) {
                  deployer.deploy(Storage);
                };
                ```

             7. 部署迁移：truffle migrate,truffle migrate会执行所有位于migrations下面的js脚本，如果该脚本在之前已经部署过一次，在没有新的迁移脚本导入的情况下，使用truffle migrate将不会两次部署该脚本，要使用truffle migrate --reset,

             8. ![1543308509380](C:\Users\ADMINI~1\AppData\Local\Temp\1543308509380.png)