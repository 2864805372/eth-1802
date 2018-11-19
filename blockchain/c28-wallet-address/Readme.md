     ## 钱包和地址
     1. 通过钱包获取地址
     2. 过程
        1. 公钥进行sha256再进行ripemd160得到公钥哈希
        2. 组成
            1. version:版本前缀，大小为1个字节，主要用来创建一个易于识别的格式
                "1"代表比特币地址
            2. pubkey hash:20个字节，公钥哈希
            3. checksum:校验和，取前4个字节，是添加到正在编码的数据一端的4个字节
                校验和通过pubkey哈希得到，用来检测输入时产生的错误
        3. version+pubkeyHash+checkSum得到两次哈希的地址
        4. 通过base58得到bitcoin地址