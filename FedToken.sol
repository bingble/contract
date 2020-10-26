pragma solidity >=0.5.16;

import './libraries/Math.sol';
import './libraries/TransferHelper.sol';

contract Fed {
    using Math for uint;

    string public name = 'FED';
    string public symbol = 'FED';
    uint8 public decimals = 18;
    uint  public totalSupply;
    mapping(address => uint) public balanceOf;
    mapping(address => mapping(address => uint)) public allowance;

    event Approval(address indexed owner, address indexed spender, uint value);
    event Transfer(address indexed from, address indexed to, uint value);
    event Deposit(address indexed from, address indexed to, uint value);
    event Withdraw(address indexed from, uint value);
    event Stopped();
    event Running();

    bool public stopped = true;
    address public owner;

    modifier isRunning() {
        require(!stopped);
        _;
    }

    modifier isStopped() {
        require(stopped);
        _;
    }

    modifier onlyFactory() {
        require(msg.sender == factory);
        _;
    }

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    bytes32 public DOMAIN_SEPARATOR;
    // keccak256("Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)");
    bytes32 public constant PERMIT_TYPEHASH = 0x6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9;
    mapping(address => uint) public nonces;

    address public factory;
    constructor(address _factory) public {
        factory = _factory;
        owner = msg.sender;

    uint chainId;
        assembly {
            chainId := chainid()
        }
        DOMAIN_SEPARATOR = keccak256(
            abi.encode(
                keccak256('EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)'),
                keccak256(bytes(name)),
                keccak256(bytes('1')),
                chainId,
                address(this)
            )
        );
    }

    function stop() onlyOwner isRunning public {
        stopped = true;
        emit Stopped();
    }

    function start() onlyOwner isStopped public {
        stopped = false;
        emit Running();
    }

    function _transfer(address from, address to, uint value) private {
        balanceOf[from] = balanceOf[from].sub(value);
        balanceOf[to] = balanceOf[to].add(value);
        emit Transfer(from, to, value);
    }

    function approve(address spender, uint value) external returns (bool) {
        _approve(msg.sender, spender, value);
        return true;
    }

    function _approve(address owner, address spender, uint value) private {
        allowance[owner][spender] = value;
        emit Approval(owner, spender, value);
    }

    function transfer(address to, uint value) external isRunning returns (bool) {
        _transfer(msg.sender, to, value);
        return true;
    }

    function transferFrom(address from, address to, uint value) external isRunning returns (bool) {
        if (allowance[from][msg.sender] != uint(- 1)) {
            allowance[from][msg.sender] = allowance[from][msg.sender].sub(value);
        }
        _transfer(from, to, value);
        return true;
    }

    function mint(address to, uint value) external onlyFactory {
        totalSupply = totalSupply.add(value);
        balanceOf[to] = balanceOf[to].add(value);
        emit Transfer(address(0), to, value);
    }

    function burn(address from, uint value) external onlyFactory {
        balanceOf[from] = balanceOf[from].sub(value);
        totalSupply = totalSupply.sub(value);
        emit Transfer(from, address(0), value);
    }

    function permit(address form, address spender, uint value, uint deadline, uint8 v, bytes32 r, bytes32 s) external {
        require(deadline >= block.timestamp, 'USDD: EXPIRED');
        bytes32 digest = keccak256(
            abi.encodePacked(
                '\x19\x01',
                DOMAIN_SEPARATOR,
                keccak256(abi.encode(PERMIT_TYPEHASH, form, spender, value, nonces[form]++, deadline))
            )
        );
        address recoveredAddress = ecrecover(digest, v, r, s);
        require(recoveredAddress != address(0) && recoveredAddress == form, 'USDD: INVALID_SIGNATURE');
        _approve(form, spender, value);
    }
}
