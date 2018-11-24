1. day09复习

   1. 以太坊背景
   2. 概念：去中化，应用开发平台
   3. 以太坊的优缺点
      1. 交易模型节省空间
      2. 开发成本降低
      3. 共识更加完善
      4. 智能合约
      5. TPS太低
   4. 钱包和浏览器
      1. etherscan.io
      2. stateoftheapps
      3. myetherwallet
   5. 智能合约
      1. 概念
      2. 优缺点
      3. 部署流程
   6. 以太坊核心概念
      1. 账户
         1. 外部账户
         2. 合约账户
      2. 账户内部结构
      3. EVM
      4. 交易
         1. 交易内容
         2. 交易吞吐量(TPS)
      5. gas(燃料)
      6. 叔区块
      7. 共识
   7. 以太坊开发环境搭建
   8. 以太坊私链搭建
      1. 创世配置genesis.json
      2. geth启动参数
      3. web3各类包与模块说明
      4. 私链搭建步骤
      5. 基本web3操作
   9. solidity基础
      1. 概念
      2. 特点
      3. 静态类型语言

2. solidity开发、测试环境搭建

   1. solidity-remix离线安装

   2. remix源码地址：https://github.com/ethereum/remix-ide

   3. 通过npm module安装

   4. ```
      npm install remix-ide -g
      remix-ide
      ```

   5. 通过源码安装

   6. ```
      git clone https://github.com/ethereum/remix-ide.git
      git clone https://github.com/ethereum/remix.git # only if you plan to link remix and remix-ide repositories and develop on it.
      cd remix-ide
      npm install
      npm run setupremix  # only if you plan to link remix and remix-ide repositories and develop on it.
      npm start
      ```

   7. VScode安装

      1. 安装vscode : https://code.visualstudio.com/download
      2. 添加solidity支持：进入扩展，搜索solidity插件，intall

   8. solidity运行环境搭建

      1. solc(solidity compiler)安装
         1. 添加ethereum源
         2. sudo apt-get update
         3. sudo apt-get install solc -y
      2. 查看solc版本
         1. solc --version

   9. solidity测试环境搭建

      1. truffle与testrpc
         1. truffle与testrpc是配套的以太坊开发框架，通过truffle可以快速编译和部署合约并且进行测试，testrpc可以快速生成以太坊测试账号
      2. 安装
         1. 安装npm
            1. 源码地址：https://nodejs.org/en/download/
            2. 下载源码包
            3. 解压tar.gz
            4. 安装
         2. 安装truffle和testrpc
            1. truffle : npm install truffle -g
            2. testrpc: npm install -g ethereumjs-testrpc
         3. 查看truffle版本
            1. truffle --version

3. solidity基础语法

   1. 通过hello world认识remix与运行环境

   2. 代码注释

      1. 单行：//
      2. 多行：/**/

   3. 文档注释

      1. 单行：///

      2. 多行：

         ```
         /**...*/
         ```

   4. 数据类型

      1. 值传递与引用传递
         1. 值传递：简单说来就是数据拷贝，对拷贝之后的数据进行操作，不会影响原数据
         2. 引用传递：修改传递之后的数据会对原始数据产生影响
      2. 值类型与引用类型
         1. 值类型
            1. 整形，布尔，地址(address)，枚举，函数，定长字节数组
            2. 在数据占用空间不大，不希望对原始数据进行修改的情况下，多数采用值类型
         2. 引用类型
            1. 不定长字节数组
            2. 字符串(string)
            3. 数组
            4. 结构体
            5. 在某些情况下，复杂类型占用的空间比较大，在拷贝的时候会占用较大的空间，所以可以考虑采用引用传递

   5. 整形

      1. 分为有符号整形和无符号整形 
      2. 在solidity中，整形有**步长**的概念，步长为8，整形范围为int8-int256,在使用整形的过程中，如果已经确定了数据大小上限的情况下，尽量使用与其最靠近的步长
      3. 整形默认是int256/uint256
      4. 在solidity中，整形支持二进制、10进制、16进制，**注意**，solidity整形不支持8进制 
      5. 支持
         1. 算术运算，比较，位运算
      6. 整形的上溢和下溢
         1. 上溢：如果一个整形变量的值达到其类型的上限，再加上一个正整数，最终结果会是变量值+正整数-类型上限
         2. 下溢：
         3. 在solidity中使用整形的时候需要注意溢出

   6. bool值

      1. 取值：true,false
      2. 支持运算符：比较，逻辑

   7. address(地址类型)

      1. 概念：代表以太坊地址，大小20字节，160位
      2. 所有地址类型变量都可以用uint160编码
      3. 地址类型支持比较运算
      4. 地址类型拥有自己的成员和属性
         1. balance:通过该属性可以获取指定地址的余额
         2. send():转账，向某个指定地址发送以太币，gas不够会导致转账失败

   8. 定长字节数组

      1. 表现形式：bytes
      2. bytes也步长，从bytes1到bytes32
      3. 默认的步长是1
      4. 支持比较运算，位运算
      5. bytes支持通过**length**获取定长字节数组的长度
      6. length属性不可修改

   9. 字符串

      1. string
      2. 对字符串的相关操作
         1. string和bytes可以转换
         2. 可以通过bytes转换之后的字符串获取长度
         3. 通过bytes转换之后可以获取字符串中指定下标的元素值
         4. 通过bytes转换之后可以修改字符串指定下标的元素值
      3. 字符串字面量不包含结束符

   10. var关键字（智能推断）

      1. 使用var 关键字的时候，在**第一次赋值**的时候，编译器会自动推断出这个变量类型，不能在函数中使用，有可能 会推断错误类型,不建议使用
      2. 在第一次变量类型已经确定之后，后面对该变量的赋值，必须在这个类型范围之内
      3. var 可以把函数值赋值给一个变量

   11. 十六进制字面量

       1. 特点：以关键hex开头，后面跟着字符串
       2. 和字符串一样，16进制也可以转换成bytes，转换完成之后可以通过下标去获取指定元素的值

   12. 常量

       1. 在solidity中，只有值类型和string支持常量 
       2. 需要注意的是，solidity中就算是定义常量，也必须把类型写在前面

   13. 数据位置(**后续补充**)

       1. 在solidity中，数据位置有三类
          1. memory:存储在内存中
          2. storage:storage存储位置修饰的变量存储的数据是永久存在的，存储在区块链上面
          3. calldata:该位置的数据是只读的，不会被持久化到区块链上，一般只有外部函数会被指定
       2. 函数的参数，返回值默认位置是memory
       3. 函数局部变量，状态变量默认存储位置是storage
       4. 数据位置转换
          1. memory->memory
          2. memory->storage
          3. storage->storage
          4. storage->memory

   14. 货币与时间单位

       1. 1eth,1finney,1szabo,1Gwei,1Mwei,1Kwei,1wei
          1. 1 eth = 10**18wei
       2. 时间单位
          1. 单位：seconds,minutes,hours,days,weeks,years

   15. 枚举

       1. 特点：
          1. 用户自定义类型
          2. 可以显示的与整形进行转换，不能隐式转换
          3. 默认从0开始
       2. 注意：枚举定义时后面不能有**";"**

   16. 结构体

       1. 结构体也是solidity中自定义数据类型

       2. 在solidity中，不能直接返回结构

       3. 初始化方式有两种

          1. 根据成员名称进行初始化
          2. 根据成员定义顺序进行初始化

       4. 两种初始化都需要添加小括号，但通过参数的定义顺序初始化不需要添加大括号

       5. ```js
          // 结构体
          contract Struct {
               // person
              struct Person {
                  string name;
                  int age;
                  string sex;
              }
              // bank
              // 嵌套
              struct Bank {
                  address addr;
                  uint balance;
                  Person p;
              }
             
              // 初始化person
              Person p1 = Person({
                  name:"troytan",
                  age:18,
                  sex:"male"
              });
              // 初始化
              Bank b1 = Bank({
                  addr:0x0,
                  balance:100,
                  p:p1
              });
          
              function getBank() view public returns(address, uint, string) {
                  return (b1.addr, b1.balance,b1.p1.name);
              }
              
              // 根据参数的定义顺序,不需要再添加"{}"
              Person p2 = Person({"troytan",20,"male"});
              function getPerson()returns(string, int, string) {
                  return(p2.name, p2.age, p2.sex);
              }
          }
          ```

   17. 数组

       1. 分类：数组分为定长数组和变长数组
       2. 声明
          1. 定长：T[K]
          2. 不定长：T[]
       3. 数组的限制
          1. 对于 storage数组来说，元素可以是任意类型
          2. 对于memory数组来说，元素的类型与函数的可见性有关，如果说函数是外部可见的，函数参数不能是一个映射类型的数组
       4. 不定长数组
          1. 可以追加元素
          2. 可以修改数组长度(注意：如果把长度减少，原有的数组中超过新的长度的元素会被截断)
       5. 数组字面量：指的是以表达式方式隐式声明一个数组，并且作为数组变量使用的一种方式

   18. 二维数组

       1. 数组元素还是一个数组
       2. 在solidity中，二维数组的行列定义与大多数语言是相反的，在插入数据是，逻辑又反过来了
       3. uint [3][4]X:代表一个有4行3列的二维数组
       4. 返回不定长二维数组时，需要添加的ABI编译
       5. ![543046922173](C:\Users\ADMINI~1\AppData\Local\Temp\1543046922173.png)

   19. 字典：键值对映射关系的存储

       1. 在mapping中，值类型可以是任意类型，键类型不能是映射，动态数组，合约，结构体，枚举几种类型
       2. solidity中mapping不支持迭代，没有长度,没有键，值集合
       3. mapping支持精准查找(通过指定的key值进行查找)，如果没有要查找的key值，返回0
       4. solidity支持delete删除mapping,但是需要注意，实际上delete并没有真正删除key值，只是将key对应的value重置为其类型的初始值

   20. 类型转换

       1. 隐式类型转换
          1. 如果运算符支持两边不同的类型，编译器会深度进行自动转换
          2. 隐式类型转换支持低位转高位，不支持高位转低位
          3. 任何一个uint160的变量都可以转换为address类型。如果该变量没有达到uint160的上限，高位补零，转换方式为将uint160的值转换为16进制
       2. 显示类型转换
          1. 在不允许隐式类型转换的情况下进行强转
          2. 如果是高位转低位，会产生截断