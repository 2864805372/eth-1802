// 支付模式之拉取

contract PullPayment {
    // 出价最高的竞标者
    address highestBidder;
    // 竞标价格
    uint highestBid;

    // 存储每个待退款的竞标者信息
    mapping(address => uint) refunds;

    function bid() {
        if(msg.value < highestBid) throw;

        if (highestBidder != 0x0) {
            refunds[highestBidder] = highestBid;
        }
        highestBidder = msg.sender;
        highestBid = msg.value;
    }
    // 实现逻辑分离，单独处理待退款的竞标者，由用户调用，而不是由合约自动推送
    function withdrawBid() {
        uint refund = refunds[msg.sender];
        refunds[msg.sender] = 0;
        if(!msg.sender.send(refund)) {
            refunds[msg.sender] = refund;
        }
    }
    
    function() payable {
        throw;
    }
}