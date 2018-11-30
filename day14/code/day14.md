1. day13复习
   1. truffle脚本命令
      1. artifacts.require
      2. exports
      3. deployer
   2. deployer api
      1. deploy
         1. 单合约构造函数带参部署
         2. 单合约构造函数无参
         3. 多个合约部署
      2. link:链接
      3. .then()
   3. 合约抽象
      1. truffle-contract
         1. 概念
         2. 安装
   4. 合约交互
      1. 分类
         1. 读取：调用
         2. 写入：交易
         3. 添加一个新的合约到网络
         4. 使用现有的合约地址生成一个新的合约实例抽象
         5. 合约充值
   5. 合约测试
      1. truffle内置编译器只能运行以.js, .es, .es6, .sol, .jsx
      2. javascript
      3. sodlity
   6. 合约调试
      1. truffle debug <transaction hash>
   7. truffle develop
      1. 内置客户端，忽略truffle.js中的配置
   8.  truffle console
      1. 没有内置客户端
   9. 指定特殊网络进行部署
      1. truffle migrate --network development
   10. truffle unbox
   11. 宠物商店 : pet-shop
2. DAPP 开发总结
   1. 以太坊开发环境
   2. slidity(合约开发)
      1. 基础语法
      2. IDE
      3. 开发框架:truffle 
      4. 编译运行环境:truffle, solc
      5. 以太坊客户端:geth
      6. 测试框架:truffle, ganache,testrpc,ganache-cli
   3. 交互实现
      1. nodejs
      2. 前端
      3. IDE
         1. ATOM,webstorm
3. ERC20,ERC721
   1. ERC20:以太坊TOKEN发布的标准接口之一
      1. 作用:为了让以太坊上各类token合约有一个共同的标准
   2. ERC721:又叫NFT，也是以太坊TOKEN发布标准接口之一(不可分隔资产)
      1. 与ERC20相比，ERC721是用于处理不可分隔资产的另一种代币标准，类似于房子，家具等
4. 高级库Openzeppelin
   1. 概念:一个帮助我们在以太坊上建立安全智能合约的开发库，当前集成在truffle开发框架中
   2. 安装
      1. truffle init
      2. npm init
      3. npm install openzeppelin-solidity
   3. 模块总述
      1. access:地址白名单和基于签名的管理
      2. crowdsale:用于管理token众筹
      3. cryptography
      4. drafts
      5. introspection
      6. lifecycle:用于管理智能合约以及其资金的生命周期的基础合约集合
      7. math:数学库
      8. ownership:用于管理合约，以及token所有权
      9. payment:管理托管，支付相关的智能合约
      10. token:一组ERC标准接口，主要就是ERC20, ERC721
      11. utils
   4. Math库详解
      1. math.sol:uint256最值与平均值获取
      2. safemath.sol:安全运算，增强合约健壮性
   5. ownership库详解
      1. Ownable.sol:合约所有权管理
   6. access库详解
      1. Roles.sol:角色管理
      2. PauserRole.sol:暂停角色管理
   7. lifecycle库详解
      1. Pauseable.sol:对合约的暂停与恢复的管理
   8. TOKEN库详解
   9. crowdsale库详解
   10. payment库详解
5. 支付模式
   1. 推送(push):合约主动维护
   2. 拉取(pull):用户主动调用