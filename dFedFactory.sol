pragma solidity >=0.5.16;

import './interfaces/IdFedBaseToken.sol';
import './interfaces/IdFedCreator.sol';
import './interfaces/IdFedPair.sol';
import './interfaces/IERC20.sol';
import './libraries/Math.sol';


contract dFedFactory {
    using Math for uint;
    mapping(address => address) public getPair;
    mapping(address => bool) public pairExist;
    address[] public allPairs;
    address public baseToken;
    address public FEDToken;
    address public creator;

    event PairCreated(address indexed token, string symbol, uint8 decimals);

    modifier onlyPair() {
        require(pairExist[msg.sender]);
        _;
    }

    address public setter;
    modifier onlySetter() {
        require(msg.sender == setter);
        _;
    }

    modifier ensure(uint _deadline) {
        require(_deadline >= block.timestamp, 'dFedFactory: EXPIRED');
        _;
    }

    constructor() public {
        setter = msg.sender;
        lastRewardedBlockGlobal = uint(- 1);
        startRewardBlock = uint(- 1);
    }

    function setBaseToken(address _baseToken) public onlySetter {
        require(baseToken == address(0), 'dFedFactory: BASE_TOKEN_HAS_BEEN_SET');
        baseToken = _baseToken;
        decimalsBaseToken = IERC20(baseToken).decimals();
    }

    function setFedToken(address _FEDToken) public onlySetter {
        require(FEDToken == address(0), 'dFedFactory: FED_TOKEN_HAS_BEEN_SET');
        FEDToken = _FEDToken;
        decimalsFED = IERC20(FEDToken).decimals();
    }

    function setCreator(address _creator) public onlySetter {
        require(creator == address(0), 'dFedFactory: FED_TOKEN_HAS_BEEN_SET');
        creator = _creator;
    }

    function setStartRewardBlock(uint _height) public onlySetter {
        startRewardBlock = _height;
        lastRewardedBlockGlobal = _height;
    }

    function allPairsLength() public view returns (uint) {
        return allPairs.length;
    }

    function createPair(address _token) external returns (address _pair) {
        require(baseToken != address(0) && _token != address(0), 'dFedFactory: ZERO_ADDRESS');
        require(getPair[_token] == address(0), 'dFedFactory: PAIR_EXISTS');
        _pair = IdFedCreator(creator).createPair(_token, baseToken);
        pairExist[_pair] = true;
        IdFedPair(_pair).initialize(baseToken, _token);
        getPair[_token] = _pair;
        allPairs.push(_pair);
        emit PairCreated(_token, IERC20(_token).symbol(), IERC20(_token).decimals());
    }

    function mintBaseToken(address _to, uint _value) external onlyPair {
        IdFedBaseToken(baseToken).mint(_to, _value);
    }

    function burnBaseToken(address _from, uint _value) external onlyPair {
        IdFedBaseToken(baseToken).burn(_from, _value);
    }

    function mintFEDToken(address _to, uint _value) external onlyPair {
        IdFedBaseToken(FEDToken).mint(_to, _value);
    }

    uint public lastRewardedBlockGlobal;
    uint public accRewardPerUSDDGlobal;
    uint public totalUSDDinLiquidityPoolGlobal;
    uint public startRewardBlock;

    function addTotalUSDDinLiquidityPoolGlobal(uint _amount) external onlyPair {
        totalUSDDinLiquidityPoolGlobal = totalUSDDinLiquidityPoolGlobal.add(_amount);
    }

    function removeTotalUSDDinLiquidityPoolGlobal(uint _amount) external onlyPair {
        totalUSDDinLiquidityPoolGlobal = totalUSDDinLiquidityPoolGlobal.sub(_amount);
    }

    uint8 decimalsFED;
    uint8 decimalsBaseToken;

    function updateInfo() external onlyPair returns (uint){
        if (block.number > lastRewardedBlockGlobal) {
            if (totalUSDDinLiquidityPoolGlobal > 0) {
                accRewardPerUSDDGlobal = accRewardPerUSDDGlobal.add((block.number.sub(lastRewardedBlockGlobal)).mul(100).mul(10 ** decimalsFED).mul(10 ** decimalsBaseToken) / (totalUSDDinLiquidityPoolGlobal));
            }
            lastRewardedBlockGlobal = block.number;
        }
        return accRewardPerUSDDGlobal;
    }
}
