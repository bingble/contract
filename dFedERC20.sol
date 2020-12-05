pragma solidity >=0.5.16;

import './libraries/Math.sol';
import './libraries/Strings.sol';
import './interfaces/IERC20.sol';

contract dFedERC20 {
    using Math for uint;
    using Strings for *;
    string public name;
    string public symbol;
    uint8 public constant decimals = 18;
    uint  public totalSupply;
    mapping(address => uint) public balanceOf;
    mapping(address => mapping(address => uint)) public allowance;

    bytes32 public DOMAIN_SEPARATOR;
    // keccak256("Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)");
    bytes32 public constant PERMIT_TYPEHASH = 0x6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9;
    mapping(address => uint) public nonces;

    event Approval(address indexed owner, address indexed spender, uint value);
    event Transfer(address indexed from, address indexed to, uint value);

    function generateLpTokenName(address _token) internal view returns (string memory) {
        return "FEDLP".toSlice().concat("-".toSlice()).toSlice().concat(IERC20(_token).symbol().toSlice());
    }

    function _mint(address _to, uint _value) internal {
        totalSupply = totalSupply.add(_value);
        balanceOf[_to] = balanceOf[_to].add(_value);
        emit Transfer(address(0), _to, _value);
    }

    function _burn(address _from, uint _value) internal {
        balanceOf[_from] = balanceOf[_from].sub(_value);
        totalSupply = totalSupply.sub(_value);
        emit Transfer(_from, address(0), _value);
    }

    function _approve(address _owner, address _spender, uint _value) private {
        allowance[_owner][_spender] = _value;
        emit Approval(_owner, _spender, _value);
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

    function permit(address _from, address _spender, uint _value, uint _deadline, uint8 _v, bytes32 _r, bytes32 _s) external {
        require(_deadline >= block.timestamp, 'dFedERC20: EXPIRED');
        bytes32 digest = keccak256(
            abi.encodePacked(
                '\x19\x01',
                DOMAIN_SEPARATOR,
                keccak256(abi.encode(PERMIT_TYPEHASH, _from, _spender, _value, nonces[_from]++, _deadline))
            )
        );
        address recoveredAddress = ecrecover(digest, _v, _r, _s);
        require(recoveredAddress != address(0) && recoveredAddress == _from, 'dFedERC20: INVALID_SIGNATURE');
        _approve(_from, _spender, _value);
    }
}
