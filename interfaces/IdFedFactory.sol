pragma solidity >=0.5.16;

interface IdFedFactory {
    event PairCreated(address indexed token, string symbol, address pair);

    function getPair(address token) external view returns (address);

    function baseToken() external view returns (address);

    function totalUSDDinLiquidityPoolGlobal() external view returns (uint);

    function accRewardPerUSDDGlobal() external view returns (uint);

    function lastRewardedBlockGlobal() external view returns (uint);

    function startRewardBlock() external view returns (uint);

    function allPairs(uint) external view returns (address);

    function setBaseToken(address _baseToken) external;

    function allPairsLength() external view returns (uint);

    function createPair(address tokenA) external returns (address);

    function mintBaseToken(address from, uint value) external;

    function burnBaseToken(address from, uint value) external;

    function mintFEDToken(address _to, uint _value) external;

    function addTotalUSDDinLiquidityPoolGlobal(uint amount) external;

    function removeTotalUSDDinLiquidityPoolGlobal(uint amount) external;

    function updateInfo() external;
}


