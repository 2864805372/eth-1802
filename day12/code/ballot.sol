pragma solidity ^0.4.18;

// 投票合约
contract Ballot{
    // 选民
    struct Voter {
        uint        weight;         // 权重
        bool        voted;          // 是否已经投过票
        uint8       vote;           // 投票对应的提案编号
        address     delegate;       // 该投票者投票权的委托对象
    }
    // 提案
    struct Proposal {
        uint voteCount;     // 该提案目前所得的票数
    }
    // 主持人
    address chairperson;

    // 提案的列表
    Proposal[]  proposals;

    // 投票地址和投票状态的对应关系
    mapping(address => Voter) voters;

    // 初始化合约，传入提案数量，以编号表示 
    function Ballot(uint8 _numProposals) public {
        chairperson = msg.sender;
        voters[chairperson].weight = 1; // 投票权重为1
        proposals.length = _numProposals; // 提案数量
    }

    // 给其它地址赋予投票权
    function giveRightToVote(address toVoter) public {
        if(msg.sender != chairperson || voters[toVoter].voted) {
            return;
        }
        voters[toVoter].weight = 1;// 初始化权重都为1
    }

    // 委托投票权
    function delegate(address to) public {
        // 获取当前地址
        Voter storage sender = voters[msg.sender];
        // 判断当前地址是否已经投过票
        if(sender.voted) return;

        while(voters[to].delegate != address(0) && voters[to].delegate != msg.sender) {
            to = voters[to].delegate; // 获取地址to所委托的地址
        }
        // 不能自己委托给自己
        if(msg.sender == to) {
            return;
        }
        // 当前地址投票状态改为已投票，意思是自己已经委托给别人了，不能再由本人来投，由地址to来投
        sender.voted = true;
        // 
        sender.delegate = to;
        Voter storage delegateTo = voters[to];
        if(delegateTo.voted) { // 如果该地址已经投过票
        // 则将当前地址投票权重添加到所投的提案编号上去
            proposals[delegateTo.vote].voteCount += sender.weight;
        } else {
            // 把当前地址的投票权重添加到委托地址的权重上去
            delegateTo.weight += sender.weight;
        }
    }

    // 根据提案编号投票
    function vote(uint8 toProposal) public {
        Voter storage sender = voters[msg.sender];
        // 如果sender已经投过票或者投票的提案编号超出了提案列表，返回
        if (sender.voted || toProposal >= proposals.length) return;
        // 表示该地址已经投过票
        sender.voted = true;
        // 投的是谁
        sender.vote = toProposal;
        proposals[toProposal].voteCount += sender.weight;
    }

    // 统计谁的票数最多
    function winningProposal() public constant returns(uint8 _winningProposal) {
        uint winningVoteCount = 0;
        // 遍历提案
        for(uint8 prop = 0; prop < proposals.length; prop++) {
            if(proposals[prop].voteCount > winningVoteCount) {
                winningVoteCount = proposals[prop].voteCount;// 替换最大值
                _winningProposal = prop; //优胜的提案编号
            }
        }
    }
}