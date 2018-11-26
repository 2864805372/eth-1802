import "./source.sol";

contract Al {
    function add(uint x, uint y) view public returns(uint) {
        return Alg.add(x, y);
    }

    function sub(uint x, uint y) view public returns(uint) {
        return Alg.sum(x, y);
    }
}
}