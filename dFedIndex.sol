pragma solidity >=0.5.16;

import './libraries/TransferHelper.sol';
import './libraries/dFedLibrary.sol';
import './interfaces/IdFedFactory.sol';
import './interfaces/IdFedPair.sol';

contract dFedIndex {
    address public factory;
    address public baseToken;

    constructor(address _factory, address _baseToken) public {
        factory = _factory;
        baseToken = _baseToken;
    }

    modifier ensure(uint _deadline) {
        require(_deadline >= block.timestamp, 'dFedIndex: EXPIRED');
        _;
    }

    function addLiquidity(
        address _token,
        uint _baseTokenAmountDesired,
        uint _tokenAmountDesired,
        uint _baseTokenAmountMin,
        uint _tokenAmountMin,
        address _to,
        uint _deadline
    ) public ensure(_deadline) returns (uint _liquidity) {
        // create the pair if it doesn't exist yet
        address _pair = IdFedFactory(factory).getPair(_token);
        if ( _pair == address(0)) {
            _pair = IdFedFactory(factory).createPair(_token);
        }

        uint _amount0;
        uint _amount1;
        (uint _reserve0, uint _reserve1)  = IdFedPair(_pair).getReserves();
        if (_reserve0 == 0 && _reserve1 == 0) {
            (_amount0, _amount1) = (_baseTokenAmountDesired, _tokenAmountDesired);
        } else {
            uint _amountTokenOptimal = dFedLibrary.quote(_baseTokenAmountDesired, _reserve0, _reserve1);
            if (_amountTokenOptimal <= _tokenAmountDesired) {
                require(_amountTokenOptimal >= _tokenAmountMin, 'dFedIndex: INSUFFICIENT_B_AMOUNT');
                (_amount0, _amount1) = (_baseTokenAmountDesired, _amountTokenOptimal);
            } else {
                uint _amountBaseTokenOptimal = dFedLibrary.quote(_tokenAmountDesired, _reserve1, _reserve0);
                assert(_amountBaseTokenOptimal <= _baseTokenAmountDesired);
                require(_amountBaseTokenOptimal >= _baseTokenAmountMin, 'dFedIndex: INSUFFICIENT_A_AMOUNT');
                (_amount0, _amount1) = (_amountBaseTokenOptimal, _tokenAmountDesired);
            }
        }

        TransferHelper.safeTransferFrom(baseToken, msg.sender, _pair, _amount0);
        TransferHelper.safeTransferFrom(_token, msg.sender, _pair, _amount1);
        _liquidity = IdFedPair(_pair).mint(_to, _amount0, _amount1);
    }

    function removeLiquidityWithPermit(
        address _token,
        uint _liquidity,
        uint _amount0Min,
        uint _amount1Min,
        address _to,
        uint _deadline,
        bool _approveMax, uint8 _v, bytes32 _r, bytes32 _s
    ) external returns (uint _amount0, uint _amount1) {
        address _pair = getPair(_token);
        uint _value = _approveMax ? uint(-1) : _liquidity;
        IdFedPair(_pair).permit(msg.sender, address(this), _value, _deadline, _v, _r, _s);
        (_amount0, _amount1) = removeLiquidity(_pair, _liquidity, _to,  _amount0Min, _amount1Min, _deadline);
    }

    function removeLiquidity(address _pair, uint _liquidity, address _to, uint _amount0Min, uint _amount1Min, uint _deadline) public ensure(_deadline) returns (uint _amount0, uint _amount1) {
        IdFedPair(_pair).transferFrom(msg.sender, _pair, _liquidity);
        (_amount0, _amount1) = IdFedPair(_pair).burn(_to);
        require(_amount0 >= _amount0Min, 'dFedIndex: INSUFFICIENT_OUTPUT_AMOUNT0');
        require(_amount1 >= _amount1Min, 'dFedIndex: INSUFFICIENT_OUTPUT_AMOUNT1');
    }

    function buyToken(
        address _token,
        uint _amountIn,
        uint _amountOutMin,
        address _to,
        uint _deadline
    ) public ensure(_deadline) {
        address _pair = getPair(_token);
        TransferHelper.safeTransferFrom(baseToken, msg.sender,_pair, _amountIn);
        IdFedPair(_pair).buyToken(_amountIn, _amountOutMin, _to);
    }

    function sellToken(
        address _token,
        uint _amountIn,
        uint _amountOutMin,
        address _to,
        uint _deadline
    ) public ensure(_deadline) {
        address _pair = getPair(_token);
        TransferHelper.safeTransferFrom(_token, msg.sender,_pair, _amountIn);
        IdFedPair(_pair).sellToken(_amountIn, _amountOutMin, _to);
    }

    function mortgage(
        address _token,
        uint _pledgeAmount,
        uint _targetNum,
        address _to,
        uint _deadline
    ) public ensure(_deadline) returns(uint _debtId) {
        address _pair = getPair(_token);
        TransferHelper.safeTransferFrom(_token, msg.sender, _pair, _pledgeAmount);
        _debtId = IdFedPair(_pair).mortgage(_pledgeAmount, _targetNum, _to);
    }

    function repay(
        address _token,
        uint _debtId,
        uint _deadline
    ) public ensure(_deadline) {
        address _pair = getPair(_token);
        uint _repayAmount = IdFedPair(_pair).getRepayById(_debtId);
        require(_repayAmount > 0, 'dFedIndex: REPAY_AMOUNT_INVALID');
        TransferHelper.safeTransferFrom(baseToken, msg.sender, _pair, _repayAmount);
        IdFedPair(_pair).repay(_debtId);
    }

    function getPair(address _token) public view returns(address _pair) {
        _pair = IdFedFactory(factory).getPair(_token);
        require( _pair != address(0), 'dFedIndex: TOKEN_PAIR_DOES_NOT_EXIST');
    }
}