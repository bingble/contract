pragma solidity >=0.5.16;

import './libraries/Math.sol';
import './libraries/TransferHelper.sol';

contract dFedUSDD {
    using Math for uint;

    string public name = 'USDD';
    string public symbol = 'USDD';
    uint8 public decimals = 6;
    uint  public totalSupply;
    mapping(address => uint) public balanceOf;
    mapping(address => mapping(address => uint)) public allowance;

    event Approval(address indexed owner, address indexed spender, uint value);
    event Transfer(address indexed from, address indexed to, uint value);
    event Deposit(address indexed from, address indexed to, uint value);
    event Withdraw(address indexed from, uint value);

    modifier onlyFactory() {
        require(msg.sender == factory);
        _;
    }

    bytes32 public DOMAIN_SEPARATOR;
    // keccak256("Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)");
    bytes32 public constant PERMIT_TYPEHASH = 0x6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9;
    mapping(address => uint) public nonces;

    address public factory;
    address public usdt;
    constructor(address _factory, address _usdt) public {
        factory = _factory;
        usdt = _usdt;

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

    function deposit(address _to, uint _amount) public {
        TransferHelper.safeTransferFrom(usdt, msg.sender, address(this), _amount);
        totalSupply = totalSupply.add(_amount);
        balanceOf[_to] = balanceOf[_to].add(_amount);
        emit Deposit(msg.sender, _to, _amount);
    }

    function withdraw(uint _amount) public {
        balanceOf[msg.sender] = balanceOf[msg.sender].sub(_amount);
        totalSupply = totalSupply.sub(_amount);
        TransferHelper.safeTransfer(usdt, msg.sender, _amount);
        emit Withdraw(msg.sender, _amount);
    }


    function _transfer(address _from, address _to, uint _value) private {
        balanceOf[_from] = balanceOf[_from].sub(_value);
        balanceOf[_to] = balanceOf[_to].add(_value);
        emit Transfer(_from, _to, _value);
    }

    function approve(address _spender, uint _value) external returns (bool) {
        _approve(msg.sender, _spender, _value);
        return true;
    }

    function _approve(address _owner, address _spender, uint _value) private {
        allowance[_owner][_spender] = _value;
        emit Approval(_owner, _spender, _value);
    }

    function transfer(address _to, uint _value) external returns (bool) {
        _transfer(msg.sender, _to, _value);
        return true;
    }

    function transferFrom(address _from, address _to, uint _value) external returns (bool) {
        if (allowance[_from][msg.sender] != uint(- 1)) {
            allowance[_from][msg.sender] = allowance[_from][msg.sender].sub(_value);
        }
        _transfer(_from, _to, _value);
        return true;
    }

    function mint(address _to, uint _value) external onlyFactory {
        totalSupply = totalSupply.add(_value);
        balanceOf[_to] = balanceOf[_to].add(_value);
        emit Transfer(address(0), _to, _value);
    }

    function burn(address _from, uint _value) external onlyFactory {
        balanceOf[_from] = balanceOf[_from].sub(_value);
        totalSupply = totalSupply.sub(_value);
        emit Transfer(_from, address(0), _value);
    }

    function permit(address _from, address _spender, uint _value, uint _deadline, uint8 _v, bytes32 _r, bytes32 _s) external {
        require(_deadline >= block.timestamp, 'USDD: EXPIRED');
        bytes32 digest = keccak256(
            abi.encodePacked(
                '\x19\x01',
                DOMAIN_SEPARATOR,
                keccak256(abi.encode(PERMIT_TYPEHASH, _from, _spender, _value, nonces[_from]++, _deadline))
            )
        );
        address recoveredAddress = ecrecover(digest, _v, _r, _s);
        require(recoveredAddress != address(0) && recoveredAddress == _from, 'USDD: INVALID_SIGNATURE');
        _approve(_from, _spender, _value);
    }
}
