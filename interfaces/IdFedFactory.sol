pragma solidity >=0.5.16;

interface IdFedFactory {
    event PairCreated(address indexed baseToken, address indexed token, address pair, uint);

    function getPair(address token) external view returns (address);

    function baseToken() external view returns (address);

    function allPairs(uint) external view returns (address);

    function setBaseToken(address _baseToken) external;

    function allPairsLength() external view returns (uint);

    function createPair(address tokenA) external returns (address);

    function mintBaseToken(address from, uint value) external ;

    function burnBaseToken(address from, uint value) external;
}


