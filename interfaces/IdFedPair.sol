pragma solidity >=0.5.16;

interface IdFedPair {
    event Approval(address indexed owner, address indexed spender, uint value);
    event Transfer(address indexed from, address indexed to, uint value);
    event DebtUpdate(address indexed owner, uint debtId, uint pledgeAmount, uint repayAmount);
    event Mint(address indexed sender, uint amount0, uint amount1);
    event Burn(address indexed sender, uint amount0, uint amount1, address indexed to);

    function name() external pure returns (string memory);

    function symbol() external pure returns (string memory);

    function decimals() external pure returns (uint8);

    function totalSupply() external view returns (uint);

    function balanceOf(address owner) external view returns (uint);

    function allowance(address owner, address spender) external view returns (uint);

    function approve(address spender, uint value) external returns (bool);

    function transfer(address to, uint value) external returns (bool);

    function transferFrom(address from, address to, uint value) external returns (bool);

    function DOMAIN_SEPARATOR() external view returns (bytes32);

    function PERMIT_TYPEHASH() external pure returns (bytes32);

    function nonces(address owner) external view returns (uint);

    function initialize(address, address, uint8) external;

    function sellToken(uint amountIn, uint amountOutMin, address to) external;

    function buyToken(uint amountIn, uint amountOutMin, address to) external;

    function mortgage(uint pledgeAmount, uint targetNum, address to) external returns (uint _debtId);

    function repay(uint debtId) external;

    function getRepayById(uint debtId) external view returns (uint repayAmount);

    function getReserves() external view returns (uint reserve0, uint reserve1);

    function mint(address to, uint amount0, uint amount1) external returns (uint liquidity);

    function burn(address to) external returns (uint amount0, uint amount1);

    function permit(address owner, address spender, uint value, uint deadline, uint8 v, bytes32 r, bytes32 s) external;

}
