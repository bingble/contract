pragma solidity >=0.5.16;

import './dFedPair.sol';

contract dFedCreator {
    modifier onlyFactory() {
        require(msg.sender == factory);
        _;
    }
    address public factory;
    constructor(address _factory) public {
        factory = _factory;
    }

    function createPair(address _token, address _baseToken) external onlyFactory returns (address _pair) {
        // single check is sufficient
        bytes memory bytecode = type(dFedPair).creationCode;
        bytes32 salt = keccak256(abi.encodePacked(_baseToken, _token));
        assembly {
            _pair := create2(0, add(bytecode, 32), mload(bytecode), salt)
        }
    }

    function getInitHash() public view returns (bytes32) {
        bytes memory bytecode = type(dFedPair).creationCode;
        bytes32 salt = keccak256(bytecode);
        return salt;
    }
}
