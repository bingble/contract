pragma solidity >=0.5.16;

import './dFedBank.sol';
import './libraries/dFedLibrary.sol';

contract dFedPair is dFedBank {
    using Math  for uint;
    constructor() public {
        debtInfos.push(debtInfo(address(0), 0, 0, 0, 0, 0, 0, 0));
        factory = msg.sender;
    }

    event BuyToken(
        address indexed sender,
        uint amountIn,
        uint amountOut,
        address indexed to
    );

    event SellToken(
        address indexed sender,
        uint amountIn,
        uint amountOut,
        address indexed to
    );

    event Mint(address indexed sender, uint amount0, uint amount1);
    event Burn(address indexed sender, uint amount0, uint amount1, address indexed to);

    uint public constant MINIMUM_LIQUIDITY = 10 ** 3;
    uint public kLast;

    function getReserves() public view returns (uint _reserve0, uint _reserve1) {
        _reserve0 = reserve0;
        _reserve1 = reserve1;
    }

    // called once by the factory at time of deployment
    function initialize(address _baseToken, address _token) external {
        require(msg.sender == factory, 'dFedPair: FORBIDDEN');
        // sufficient check
        token0 = _baseToken;
        token1 = _token;

        name = generateLpTokenName(_token);
        symbol = name;
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

    function sellToken(uint _amountIn, uint _amountOutMin, address _to) external lock {
        require(_amountIn > 0, 'dFedPair: INVALID_INPUT');
        require(IERC20(token1).balanceOf(address(this)).sub(reserve1).sub(totalPledge).sub(totalRefundToken1Amount) >= _amountIn, 'dFedPair: INSUFFICIENT_INPUT_AMOUNT');
        require(_amountOutMin < reserve0, 'dFedPair: INSUFFICIENT_LIQUIDITY');
        (uint _actAmountIn, uint _debtAmount) = liquidate(_amountIn);
        uint _totalAmountOut = dFedLibrary.getAmountOut(_actAmountIn, reserve1, reserve0);
        uint _amountOut = _totalAmountOut.sub(_debtAmount);
        uint _fee = _amountOut.mul(3) / 1000;
        uint _actAmountOut = _amountOut.sub(_fee);
        require(_actAmountOut >= _amountOutMin, 'dFedPair: INSUFFICIENT_OUTPUT_AMOUNT');
        if (_actAmountOut > 0) TransferHelper.safeTransfer(token0, _to, _actAmountOut);
        if (_fee > 0) fee = fee.add(_fee);
        updateReserves(reserve0.sub(_totalAmountOut), reserve1.add(_actAmountIn));
        emit SellToken(msg.sender, _amountIn, _actAmountOut, _to);
    }

    function buyToken(uint _amountIn, uint _amountOutMin, address _to) external lock {
        require(_amountIn > 0, 'dFedPair: INVALID_INPUT');
        require(IERC20(token0).balanceOf(address(this)).sub(reserve0).sub(fee).sub(totalRefundToken0Amount) >= _amountIn, 'dFedPair: INSUFFICIENT_INPUT_AMOUNT');
        uint _fee = _amountIn.mul(3) / 1000;
        uint _actAmountIn = _amountIn.sub(_fee);
        uint _amountOut = dFedLibrary.getAmountOut(_actAmountIn, reserve0, reserve1);
        require(_amountOut >= _amountOutMin, 'dFedPair: INSUFFICIENT_OUTPUT_AMOUNT');
        if (_amountOut > 0) TransferHelper.safeTransfer(token1, _to, _amountOut);
        if (_fee > 0) fee = fee.add(_fee);
        updateReserves(reserve0.add(_actAmountIn), reserve1.sub(_amountOut));
        emit BuyToken(msg.sender, _amountIn, _amountOut, _to);
    }

    function updateReserves(uint _reserve0, uint _reserve1) private {
        require(_reserve0.mul(_reserve1) >= uint(reserve0).mul(reserve1), 'dFedPair: K');
        reserve0 = _reserve0;
        reserve1 = _reserve1;
    }

    // this low-level function should be called from a contract which performs important safety checks
    function mint(address _to, uint _amount0, uint _amount1) external lock returns (uint _liquidity) {
        require(_amount0 <= IERC20(token0).balanceOf(address(this)).sub(reserve0).sub(fee), 'dFedPair: INSUFFICIENT_AMOUNT_0');
        require(_amount1 <= IERC20(token1).balanceOf(address(this)).sub(reserve1).sub(totalPledge), 'dFedPair: INSUFFICIENT_AMOUNT_1');

        uint _totalSupply = totalSupply;
        // gas savings, must be defined here since totalSupply can update in _mintFee
        if (_totalSupply == 0) {
            _liquidity = Math.sqrt(_amount0.mul(_amount1)).sub(MINIMUM_LIQUIDITY);
            _mint(address(0), MINIMUM_LIQUIDITY);
            // permanently lock the first MINIMUM_LIQUIDITY tokens
        } else {
            _liquidity = Math.min(_amount0.mul(_totalSupply) / reserve0, _amount1.mul(_totalSupply) / reserve1);
        }
        require(_liquidity > 0, 'dFedPair: INSUFFICIENT_LIQUIDITY_MINTED');
        _mint(_to, _liquidity);

        require((reserve0.add(_amount0)).mul(reserve1) >= (reserve1.add(_amount1).mul(reserve0)), 'dFedPair: INVALID_UPDATE');
        reserve0 = reserve0.add(_amount0);
        reserve1 = reserve1.add(_amount1);
        // reserve0 and reserve1 are up-to-date
        emit Mint(_to, _amount0, _amount1);
    }

    // this low-level function should be called from a contract which performs important safety checks
    function burn(address _to) external lock returns (uint _amount0, uint _amount1) {
        uint liquidity = balanceOf[address(this)];
        uint _totalSupply = totalSupply;
        // gas savings, must be defined here since totalSupply can update in _mintFee
        _amount0 = liquidity.mul(reserve0) / _totalSupply;
        // using balances ensures pro-rata distribution
        _amount1 = liquidity.mul(reserve1) / _totalSupply;
        // using balances ensures pro-rata distribution
        require(_amount0 > 0 && _amount1 > 0, 'dFedPair: INSUFFICIENT_LIQUIDITY_BURNED');
        _burn(address(this), liquidity);
        (uint _getToken1Amount, uint _payToken0Amount, uint _refundToken0Amount) = overlapDebit(reserve0.sub(_amount0), reserve1.sub(_amount1));
        refundAmount[_to] = refundAmount[_to].add(_refundToken0Amount);
        reserve0 = reserve0.sub(_amount0);
        reserve1 = reserve1.sub(_amount1);
        _amount1 = _amount1.add(_getToken1Amount);
        _amount0 = _amount0.sub(_payToken0Amount);
        if (_amount1 > 0) TransferHelper.safeTransfer(token1, _to, _amount1);
        if (_amount0 > 0) TransferHelper.safeTransfer(token0, _to, _amount0);
        // reserve0 and reserve1 are up-to-date
        emit Burn(msg.sender, _amount0, _amount1, _to);
    }

    function withdrawToken0(address _to, uint _amount) public {
        require(totalRefundToken0Amount > _amount, 'dFedPair: INSUFFICIENT_REFUND_BURNED');
        refundAmount[_to] = refundAmount[_to].sub(_amount);
        if (_amount > 0) TransferHelper.safeTransfer(token0, _to, _amount);
    }

    function withdrawToken1(address _to, uint _amount) public {
        require(totalRefundToken1EqualToken0Amount > _amount, 'dFedPair: INSUFFICIENT_REFUND_BURNED');
        refundAmount[_to] = refundAmount[_to].sub(_amount);
        uint _amount1 = _amount.mul(totalRefundToken1EqualToken0Amount) / totalRefundToken1Amount;
        if (_amount1 > 0) TransferHelper.safeTransfer(token1, _to, _amount1);
    }
}