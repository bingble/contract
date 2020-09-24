pragma solidity >=0.5.16;

import './libraries/Math.sol';
import './libraries/TransferHelper.sol';

contract dFedUSDD {
    using Math for uint;

    string public name = 'USDD';
    string public symbol = 'USDD';
    uint8 public decimals = 18;
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

    address public factory;
    address public usdt;
    constructor(address _factory, address _usdt) public {
        factory = _factory;
        usdt = _usdt;
    }

    function deposit(address to, uint amount) public {
        TransferHelper.safeTransferFrom(usdt, msg.sender, address(this), amount);
        totalSupply = totalSupply.add(amount);
        balanceOf[to] = balanceOf[to].add(amount);
        emit Deposit(msg.sender, to, amount);
    }

    function withdraw(uint amount) public {
        balanceOf[msg.sender] = balanceOf[msg.sender].sub(amount);
        totalSupply = totalSupply.sub(amount);
        TransferHelper.safeTransfer(usdt, msg.sender, amount);
        emit Withdraw(msg.sender, amount);
    }


    function _transfer(address from, address to, uint value) private {
        balanceOf[from] = balanceOf[from].sub(value);
        balanceOf[to] = balanceOf[to].add(value);
        emit Transfer(from, to, value);
    }

    function approve(address spender, uint value) external returns (bool) {
        allowance[msg.sender][spender] = value;
        emit Approval(msg.sender, spender, value);
        return true;
    }

    function transfer(address to, uint value) external returns (bool) {
        _transfer(msg.sender, to, value);
        return true;
    }

    function transferFrom(address from, address to, uint value) external returns (bool) {
        if (allowance[from][msg.sender] != uint(- 1)) {
            allowance[from][msg.sender] = allowance[from][msg.sender].sub(value);
        }
        _transfer(from, to, value);
        return true;
    }

    function mint(address to, uint value)  external onlyFactory {
        totalSupply = totalSupply.add(value);
        balanceOf[to] = balanceOf[to].add(value);
        emit Transfer(address(0), to, value);
    }

    function burn(address from, uint value)  external onlyFactory {
        balanceOf[from] = balanceOf[from].sub(value);
        totalSupply = totalSupply.sub(value);
        emit Transfer(from, address(0), value);
    }
}
