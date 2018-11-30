// 支付模式之推送
// 实现一个简单竞标合约
/*
    在竞标过程中，如果有出价更高的人，会将之前竞标成功的人的钱退回
    但是，如果第一个人在自己的回退函数中植入恶意代码，使得send函数一直返回false，
    就无法更新竞标成功者的信息，使得自己以低价竞标成功
 */
contract PushPayments {
    // 出价最高的竞标者
    address highestBidder;
    // 竞标价格
    uint highestBid;
    
    // 回退函数
    function() payable{
         throw;
    }
    // 竞标函数
    function bid() {
        // 如果当前的价格小于竞标的价格，抛出
        if(msg.value <= highestBid) throw;
        if(highestBidder != 0x0) {
            if(!highestBidder.send(highestBid)) {
                throw;
            }
        }
        highestBidder = msg.sender;
        highestBid = msg.value;
    }
}
