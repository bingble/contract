pragma solidity >=0.5.16;

interface IdFedCreator {
    function createPair(address token, address baseToken) external returns (address);
}
