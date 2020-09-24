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

    event PairCreated(address indexed token0, address indexed token1, address pair, uint);

    modifier onlyPair() {
        require(pairExist[msg.sender]);
        _;
    }

    address public baseTokenSetter;
    modifier onlyBaseTokenSetter() {
        require(msg.sender == baseTokenSetter);
        _;
    }

    modifier ensure(uint _deadline) {
        require(_deadline >= block.timestamp, 'dFedFactory: EXPIRED');
        _;
    }

    constructor() public {
        baseTokenSetter = msg.sender;
    }

    function setBaseToken(address _baseToken) public onlyBaseTokenSetter {
        baseToken = _baseToken;
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
        dFedPair(_pair).initialize(baseToken, _token);
        getPair[_token] = _pair;
        pairExist[_pair] = true;
        allPairs.push(_pair);
        emit PairCreated(baseToken, _token, _pair, allPairs.length);
    }

    function mintBaseToken(address _to, uint _value) external onlyPair {
        IdFedBaseToken(baseToken).mint(_to, _value);
    }

    function burnBaseToken(address _from, uint _value) external onlyPair {
        IdFedBaseToken(baseToken).burn(_from, _value);
    }
}
