pragma solidity >=0.5.16;

import './interfaces/IdFedBaseToken.sol';
import './dFedPair.sol';
import './libraries/Math.sol';


contract dFedFactory {
    using Math for uint;
    mapping(address => address) public getPair;
    mapping(address => bool) public pairExist;
    address[] public allPairs;
    address public baseToken;
    address public FEDToken;

    event PairCreated(address indexed token, string symbol, address pair);

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
        // single check is sufficient
        bytes memory bytecode = type(dFedPair).creationCode;
        bytes32 salt = keccak256(abi.encodePacked(baseToken, _token));
        assembly {
            _pair := create2(0, add(bytecode, 32), mload(bytecode), salt)
        }

        dFedPair(_pair).initialize(baseToken, _token, decimalsFED);
        getPair[_token] = _pair;
        pairExist[_pair] = true;
        allPairs.push(_pair);
        emit PairCreated(_token, IERC20(_token).symbol(), _pair);
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

    function updateInfo() external onlyPair {
        if (block.number > lastRewardedBlockGlobal && totalUSDDinLiquidityPoolGlobal > 0) {
            accRewardPerUSDDGlobal = (block.number.sub(lastRewardedBlockGlobal)).mul(100).mul(10 ** decimalsFED).mul(10 ** decimalsBaseToken) / (totalUSDDinLiquidityPoolGlobal);
            lastRewardedBlockGlobal = block.number;
        }
    }
}
