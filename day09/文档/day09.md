1. day08复习

   1. nodejs：运行在服务端的js，基于v8引擎
   2. nodejs功能：实现高并发的web服务
   3. 特点：异步I/O，事件驱动
   4. 下载安装
   5. 交互式控制台
   6. 基本的nodejs应用
   7. 包管理工具npm
   8. npm使用
   9. cnpm
   10. 如何在repl上面编写和运行nodejs程序
       1. .save filename
       2. .load filename
   11. nodejs变量
       1. var 
       2. let
   12. nodejs函数
       1. 函数可以作为作为参数传递
       2. es6:箭头函数=>
   13. nodejs回调函数
   14. nodejs异常处理
       1. try catch
       2. call back
       3. events
       4. promise
       5. domain
   15. 事件-->EventEmitter
   16. stream
   17. 模块系统
       1. 创建模块：exports
       2. 引用模块：require
       3. 模块查找步骤
   18. 全局对象：可以在程序的任何地方使用
   19. 常用模块
       1. util:提供常用函数的集合
       2. path:处理文件路径
   20. node get/post请求
   21. express：nodejs的web应用框架
       1. 核心特性

2. 以太坊

   1. 以太坊产生背景

      1. 为了解决比特币只能应用于支付领域而设计的一种新的公链

   2. 概念

      1. 以太坊是一个建立在区块技术基础之上的去中心化应用开发平台。允许任何人在该平台中建立和使用通过区块链技术产生的去中化应用(DAPPS)。被区块链2.0

   3. 优势

      1. 和比特币相比，以太坊对底层区块链技术做了一个完整的封装，使得区块链开发者可以直面上层应用平台开发，只需要专注于应用本身，降低了开发难度和成本
      2. 与bitcoin相比，以太坊采用的是账户模型，不再使用UTXO
         1. 节省空间，每笔交易只有一个输入，只有一个输出，一个签名
         2. 更加容易理解
      3. 共识算法
         1. 以太坊采用的是改进之后的POW，避免了比特币挖的能耗问题
         2. ethash：内存难解，不再是单纯的依靠算力
      4. 智能合约：和比特币相比，以太坊提供了一个图灵完备的开发语言-solidity

   4. 缺陷

      1. TPS低，不适合大量快速交易
      2. 以太坊网络节点不适合存储大文件

   5. 以太币

      1. 以太坊内置货币叫做以太币(ETH/ETC)
      2. 发行量
         1. 众筹其间发行7200W
         2. 以后每年发行1800W
      3. 区块奖励
         1. 以太币每产生一个新的区块奖励5个以太币，以太币每小时产生250个区块
      4. 叔块奖励
         1. 如果该叔块被引用，那么挖出该块的矿工有大约4.375个以太币奖励
      5. 叔块引用奖励
         1. 矿工每引用一个叔块，可以得到大约0.15个以太币奖励，最多引用两个叔块
      6. 单位：以太坊最小单位：wei。 1 eth = 10^18 wei

   6. 以太坊钱包

      1. 钱包介绍：以太坊钱包就是一个以太客户端，可以理解为一个开发者工具，主要提供账户管理，转账，挖矿，智能合约部署
      2. 分类
         1. 交互式控制台，主要为geth
         2. 图形化钱包
            1. 浏览器钱包:MetaMask
            2. 客户端钱包:mist
            3. 手机端钱包:imtoken

   7. 以太坊数据浏览器

      1. 区块数据浏览器:etherscan.io，可以对交易，区块，数据等查询
      2. DAPP浏览器:stateofthedapps,浏览以太坊DAPP的网站
      3. 浏览器wallet:myetherwallet.com:通过该网站可以实现对以太坊账户管理

   8. 智能合约

      1. 概念：在计算机领域，智能合约指一种计算机协议，该协议拥有自称校验、自称执行的功能，不需要人为干预
      2. 优点
         1. 精确执行
         2. 较低的人为干预
         3. 去中心化
      3. 缺点
         1. 合约一旦部署，将不接受修改，如果合约本身出现问题，将会无法修复
      4. 合约部署流程
         1. 启动一个以太坊节点
         2. 编写智能合约
         3. 编译
         4. 部署到以太坊上面 
         5. 可以使用web3.js等库提供的API进行合约调用 
      5. 以太坊货币单位
         1. eth
         2. finney
         3. gwei
         4. wei

   9. 以太坊核心概念

      1. EVM-以太坊虚拟机：以太坊中智能合约的运行环境

      2. 账户模型

         1. 外部账户：与比特币钱包账户相似，被公钥-私钥对控制

         2. 合约账户：被存储在账户中的合约代码控制

         3. 区别

         4. |                    | 外部账户    | 合约账户     |
            | ------------------ | ----------- | ------------ |
            | 是否可以拥有以太币 | 是          | 是           |
            | 是否可以发送交易   | 是          | 是           |
            | 账户由谁控制       | 公钥-私钥对 | 合约代码     |
            | 是否包含代码       | 不包含代码  | 包含合约代码 |

         5. 账户结构

            1. nonce、balance、storeageRoot、codeHash

            2. |             | 外部账户                                                    | 合约账户                               |
               | ----------- | ----------------------------------------------------------- | -------------------------------------- |
               | nonce       | 代表该账户发送过的交易数量，在以坊中，可以通过nonce防止双花 | 代表该账户创建过的合约数量             |
               | balance     | 账户余额                                                    | 账户余额                               |
               | storageRoot | merkle根节点，默认为空                                      | merkle根节点，默认为空                 |
               | codeHash    | 空哈希                                                      | 账户的EVM code(编译后的智能合约字节码) |

         6. 交易：从外部账户发出消息(签名的数据包)

            1. 交易内容
               1. 消息接收者
               2. 用于确认发送的签名
               3. 账户余额
               4. 发送的额外数据
               5. Gas

         7. TPS(交易吞吐量)

            1. 交易吞吐量就是批以太坊每秒能够处理的交易数量
            2. 计算方法：TPS=gaslimit/gas/出块速度

         8. GAS

            1. gas就是我们所说的燃料(手续费)，以太坊的每一笔交易都会收取一定数量的gas,设置gas的目的就是限制交易步数，同时为交易的执行支付费用 
            2. 在执行交易的过程中，如果gas消耗完了，但交易还没有完成执行，会触发out-of-gas，同时当前调用所修改的所有状态都会被回滚。被消耗的gas不会退还
            3. 如果在交易执行完成之后，消耗的gas小于设置的gas上限，多余gas会退还到原地址(交易的发起者)
            4. gasLimit:在交易过程中最多允许消耗的gas数量，由交易发起人来决定
            5. gasUsed:交易过程中实际消耗的gas
            6. gasPrice:gas单价，以gwei表示，gasprice是打包交易时选取交易的重要衡量标准

         9. 以太坊状态转换：交易执行过程中，以太坊从一个状态转换到另一个状态的过程 

         10. 共识

             1. 以太坊当前仍然采用POW

         11. 常见的DAPP

             1. CK(cryptoKitties)
             2. DAO:去中心化的自治组织

         12. 叔区块

             1. 如果一个块不是最长链的一部分，那么它被称为孤块，一个孤块也是一个合法的区块
             2. 为什么以太坊要引用叔区块
                1. 以太坊十几秒的出场时间，大大的增加了孤块的产生，由于孤块本身也是合法区块，所以引用叔区块进行验证，可以使得主链获得更多的安全性保障
             3. 一次最多只能引用两个叔块

      3. 以太坊开发环境搭建

         1. git安装
            1. windows:https://www.git-scm.com/download/win
         2. 安装geth
            1. 快速安装：https://geth.ethereum.org/downloads/ 下载对应版本
         3. 通过源码安装
            1. git下载源码：git clone https://github.com/ethereum/go-ethereum
            2. 进入源码目录
            3. make geth或者make all

      4. 以太坊搭建私链

         1. 新建一个区块链目录
         2. 生成创世区块
            1. 在目录中新建一个配置文件genesis.json，存储创世配置
            2. geth --datadir data init genesis.json
         3. 启动节点：geth --datadir data console
            1. 节点启动之后，会在data下生成geth和keystore
               1. geth:存储区块链区块数据
               2. keystore:存储私钥文件

      5. 创世区块配置文件参数解析

         - chainid:独立的区块链网络ID，公网的ID是1，注意：不同网络ID之间的节点无法互相连接
         - alloc:预置账号以及账号的以太币数量，私链用不上
         - coinbase:矿工账号，在私链中，默认会把账号列表中的第一个账号当成coinbase
         - difficulty:当前区块挖矿难度，可以动态调节
         - extraData:附加信息，随便填
         - gaslimit:gas上限
         - nonce:随机数，用于挖矿
         - mixhash:和nonce结合起来用的
         - parentHash:前区块哈希
         - timeStamp:创世区块时间戳

      6. geth启动参数详解

         - --datadir:指定区块链区块数据存储位置
         - --rpc:启动RPC通信，可以进行智能合约部署和调度
         - --rpcaddr=0.0.0.0:指定HTTP-RPC监听地址，默认localhost
         - --rpcport=8545:指定HTTP-RPC监听端口，默认8545
         - --rpccorsdomain "*":允许跨域请求的域名列表(逗号分隔)
         - --rpcapi:允许连接的RPC客户端
         - --ws:启动websocket通信
         - --nodiscover:关闭自动连接(不去发现其它节点)
         - --maxpeers 30:允许的最大连接数，默认是25个
            1. --mine:开启挖矿，默认CPU挖矿
         - --minerthreads 1:用于挖矿的CPU线程数，默认是1
         - --etherbase:矿工账号

      7. geth启动指令

         1. console:启动交互式命令行，可以在geth中执行命令

      8. geth模块对象说明

         - admin:提供管理节点的相关方法
         - eth:提供操作区块链的相关方法
         - miner:提供启动挖矿和停止挖矿的
         - net:提供查找P2P网络状态的相关方法
         - personal:提供管理账户的相关方法
         - rpc:提供HTTP-RPC访问的相关方法
         - txpool:提供查找交易内存池的方法
         - web3:包含以上所有对象的方法，还提供了一些单位换算的方法

      9. 基本操作

         1. 查看账户列表：eth.accounts
         2. 账号管理
            1. 新建账号
               1. personal.newAccount("123456")
               2. personal.newAccount()，enter再输入密码
            2. 账号解锁
               1. personal.unlockAccount(account)
            3. 挖矿
               1. 设置挖矿地址：miner.setEtherbase(account)
               2. 查看挖矿地址：eth.coinbase
               3. 启动挖矿：miner.start()
               4. 停止挖矿：miner.stop()
            4. 转账：
               1. eth.sendTransaction({from:"from",to:"to",value:web3.toWei(1,"ether")})
               2. 必须处于挖矿的状态
               3. **如果是通过geth启动的私链节点，必须要添加上--networkid private_networkid,否则当前节点默认仍然是主网的网络ID**
            5. 查看交易池:txpool.status
            6. 查看pending交易详情:txpool.inspect.pending
            7. 查询余额:eth.getBalance(account)
            8. 针对交易的操作
               1. 查看发起交易时的详情:eth.getTransaction("txHash")
            9. 针对区块的操作
               1. 查看当前区块总数： eth.blockNumber
               2. 查看最新的区块信息：eth.getBlock('latest')
               3. 查看指定区块信息：eth.getBlock('number')
            10. 远程节点管理
               1. 查看节点信息：admin.nodeInfo
               2. 获取节点名称：admin.nodeInfo.enode
               3. 添加其它节点：admin.addPeer(node_name)
               4. 查看已连接的远程节点：admin.peers

      10. 以太坊联盟链搭建(**作业**)

         1. 在联盟链中，所有节点的创世区块配置文件必须使用同一份genesis.json

         2. 创建联盟链账户

         3. 使用genesis.json在每台机器上面创建联盟链节点

         4. 搭建联盟链网络

            1. 在每台节点上启动geth

            2. 获取节点地址

            3. ```
               "enode://1d85054b86443266af59330673c6bb82403a65e82f69e4244b3beac9ea77f3b69d8a96382ebcefa17bf909fceaa3f9fd96f4a618eb147cfd1c2cf7effd72cfc6@[::]:30303"
               ```

            4. 把地址信息"[::]"替换成该节点的公网IP

            5. 在每一个节点的data/geth目录下创建一个静态节点文件static-nodes.json,将其它节点的enode信息写入

            6. 启动之后，通过admin.peers查看

      11. solidity开发语言

          1. 概念：一种用于开发智能合约的高级语言，用于编写以太坊智能合约，在EVM上运行

          2. 语法类似于javascript，是一种面向对象语言

          3. 语法特点(和传统语言的区别)

             1. Address类型：由于以太坊底层是基于账户的，主要作用是定位合约，账户，合约代码
             2. payable关键字：鉴于以太坊本身的支付属性，在内部框架中支持支付，payable关键字可以让我们直接在语言层面支持支付
             3. 可见性：除去传统语言中所拥有的publice,private，还支持external, internal
             4. 数据位置分类：与传统语言不同，solidity分为状态变量和内存变量，状态变量永久存在(保存在合约的存储空间中)
             5. 异常机制：在solidity中，一旦出现异常，所有执行都会被回滚，主要是为了保证合约执行的原子性，不允许中间状态的存在

          4. solidity是一种静态类型的语言,在编译时就需要明确指定变量类型

          5. solidityIDE

             1. remix:在线solidity编译器:https://remix.ethereum.org/#optimize=false&version=soljson-v0.4.24+commit.e67f0147.js

             