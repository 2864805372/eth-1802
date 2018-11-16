1. day03复习
   1. 比特币交易原理
      1. 与传统web交易的区别
      2. 特性：堆栈式脚本语言，没有余额，账户的概念
      3. UTXO交易模式
         1. 输入
         2. 输出
      4. UTXO查找方式
         1. 如何计算指定账户的UTXO集合
         2. 可以将UTXO理解一个币
      5. 如何判断一笔交易是否是有效交易
         1. 有效签名
         2. 有足够金额
      6. 交易过程
         1. 保证用户只会使用属于自己的比特币
         2. 保证交易有效
         3. 交易分类
            1. coinbase:没有输入
            2. 普通转账:正常情况下的转账交易，有输入
         4. UTXO**不可分割**，只要使用一次就不能再使用
            1. 设计UTXO的目的，为了防止双花
      7. 真实的比特币交易
         1. 真实交易的组成
         2. 输入、输出脚本
      8. 命令行实现
         1. os.args
         2. flags
      9. 获取区块链对象
      10. 交易相关基础
          1. 交易结构
          2. coinbase交易
      11. 通过命令行实现发起转账的初步功能
      12. 通过发起转账实现挖，生成新的区块