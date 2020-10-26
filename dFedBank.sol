pragma solidity >=0.5.16;

import './libraries/TransferHelper.sol';
import './dFedERC20.sol';
import './interfaces/IdFedFactory.sol';

contract dFedBank is dFedERC20 {
    debtInfo[] public debtInfos;
    using Math for uint;

    //    event Mortgage(address indexed  sender, uint pledgeAmount, uint targetAmount);
    //    event Repay(address indexed sender, uint debtId, uint pledgeAmount, uint repayAmount);
    //    event Liquidate(address indexed sender, uint debtId, uint pledgeAmount, uint repayAmount);
    event DebtUpdate(address indexed owner, uint debtId, uint pledgeAmount, uint repayAmount);

    bool private unlocked = true;
    uint public fee;
    modifier lock() {
        require(unlocked, 'dFedBank: LOCKED');
        unlocked = false;
        _;
        unlocked = true;
    }

    struct debtInfo {
        address user;
        uint debtId;
        uint pledgeAmount;
        uint repayAmount;
        uint debtToken0Amount;
        uint debtToken1Amount;
        uint nextIndex;
        uint lastIndex;
    }

    address public token0;
    address public token1;
    address public factory;

    uint public reserve0;
    uint public reserve1;

    uint public totalPledge;
    uint public uniqueDebtId = 1;
    uint lastNodeIndex;

    uint public totalRefundToken0Amount;
    uint public totalRefundToken1Amount;
    uint public totalRefundToken1EqualToken0Amount;

    mapping(address => uint) public refundAmount;

    function mortgage(uint _pledgeAmount, uint _targetNum, address _to) external lock returns (uint _debtId){
        require(IERC20(token1).balanceOf(address(this)).sub(reserve1).sub(totalPledge).sub(totalRefundToken1Amount) >= _pledgeAmount, 'dFedBank: INSUFFICIENT_INPUT_AMOUNT');
        require(_targetNum > 0, 'dFedBank: TARGET_NUMBER_INVALID');
        (uint _index, bool _valid) = isValidDebit(_pledgeAmount, _targetNum);
        require(_valid, 'dFedBank: INVALID_DEBIT');
        _debtId = insertDebit(_index, _pledgeAmount, _targetNum, _to);
        IdFedFactory(factory).mintBaseToken(_to, _targetNum);
        // emit Mortgage(msg.sender, _pledgeAmount, _targetNum);
        emit DebtUpdate(_to, _debtId, _pledgeAmount, _targetNum);
    }

    function isValidDebit(uint _pledgeAmount, uint _targetNum) public view returns (uint _index, bool _valid) {
        uint _x = getStartPoint(_pledgeAmount, _targetNum, reserve0, reserve1);
        if (_x < reserve1) {
            return (0, false);
        }
        uint _tmpIndex;
        uint _sellPoint;
        for (; debtInfos[_index].nextIndex != 0; _index = debtInfos[_index].nextIndex) {
            _tmpIndex = debtInfos[_index].nextIndex;
            _sellPoint = getStartPoint(debtInfos[_tmpIndex].pledgeAmount, debtInfos[_tmpIndex].repayAmount, reserve0, reserve1);
            if (_x.add(_pledgeAmount) <= _sellPoint) {
                return (_index, true);
            }

            if (_x < _sellPoint.add(debtInfos[_tmpIndex].pledgeAmount)) {
                return (0, false);
            }
        }
        return (_index, true);
    }

    function repay(uint _debtId) external lock {
        uint _index = getDebtIndexById(_debtId);
        require(_index != 0, 'dFedBank: DEBIT_NOT_FOUND');

        uint _tmpRepayAmount = debtInfos[_index].repayAmount.add(debtInfos[_index].debtToken0Amount);
        uint _tmpPledgeAmount = debtInfos[_index].pledgeAmount.add(debtInfos[_index].debtToken1Amount);
        address _user = debtInfos[_index].user;
        require(IERC20(token0).balanceOf(address(this)).sub(reserve0).sub(fee).sub(totalRefundToken0Amount) >= _tmpRepayAmount, 'dFedBank: INSUFFICIENT_INPUT_AMOUNT');
        totalRefundToken0Amount = totalRefundToken0Amount.add(debtInfos[_index].debtToken0Amount);
        delDebit(_index);
        IdFedFactory(factory).burnBaseToken(address(this), _tmpRepayAmount);
        TransferHelper.safeTransfer(token1, _user, _tmpPledgeAmount);
        //emit Repay(msg.sender, _debtId, _tmpPledgeAmount, _tmpRepayAmount);
        emit DebtUpdate(_user, _debtId, 0, 0);
    }

    function getDebtIndexById(uint _debtId) public view returns (uint _index) {
        uint _tmpIndex = debtInfos[0].nextIndex;
        while (_tmpIndex != 0) {
            if (debtInfos[_tmpIndex].debtId == _debtId) {
                return _tmpIndex;
            }
            _tmpIndex = debtInfos[_tmpIndex].nextIndex;
        }
        return 0;
    }

    function getRepayById(uint _debtId) public view returns (uint _repayAmount) {
        uint _index = getDebtIndexById(_debtId);
        require(_index != 0, 'dFedBank: DEBIT_NOT_FOUND');
        _repayAmount = debtInfos[_index].repayAmount.add(debtInfos[_index].debtToken0Amount);
    }

    function liquidate(uint _amountIn) internal returns (uint _actAmountIn, uint _debtAmount) {
        _actAmountIn = _amountIn;
        uint _tmpIndex = debtInfos[0].nextIndex;
        while (_tmpIndex != 0) {
            if (debtInfos[_tmpIndex].pledgeAmount == 0 && debtInfos[_tmpIndex].repayAmount == 0) {
                if (getStartPoint(debtInfos[_tmpIndex].debtToken1Amount, debtInfos[_tmpIndex].debtToken0Amount, reserve0, reserve1) < reserve1.add(_actAmountIn)) {
                    uint _nextIndex = debtInfos[_tmpIndex].nextIndex;
                    totalRefundToken1Amount = totalRefundToken1Amount.add(debtInfos[_tmpIndex].debtToken1Amount);
                    totalRefundToken1EqualToken0Amount = totalRefundToken1EqualToken0Amount.add(debtInfos[_tmpIndex].debtToken0Amount);
                    delDebit(_tmpIndex);
                    _tmpIndex = _nextIndex;
                    continue;
                } else {
                    break;
                }
            }

            if (getStartPoint(debtInfos[_tmpIndex].pledgeAmount, debtInfos[_tmpIndex].repayAmount, reserve0, reserve1) < reserve1.add(_actAmountIn)) {
                uint _nextIndex = debtInfos[_tmpIndex].nextIndex;
                _actAmountIn = _actAmountIn.add(debtInfos[_tmpIndex].pledgeAmount);
                _debtAmount = _debtAmount.add(debtInfos[_tmpIndex].repayAmount);
                // emit Liquidate(debtInfos[_tmpIndex].user, debtInfos[_tmpIndex].debtId,debtInfos[_tmpIndex].pledgeAmount,debtInfos[_tmpIndex].repayAmount);
                emit DebtUpdate(debtInfos[_tmpIndex].user, debtInfos[_tmpIndex].debtId, 0, 0);
                totalRefundToken1Amount = totalRefundToken1Amount.add(debtInfos[_tmpIndex].debtToken1Amount);
                totalRefundToken1EqualToken0Amount = totalRefundToken1EqualToken0Amount.add(debtInfos[_tmpIndex].debtToken0Amount);
                delDebit(_tmpIndex);
                _tmpIndex = _nextIndex;
            } else {
                break;
            }
        }
    }

    function overlapDebit(uint _reserve0, uint _reserve1) internal returns (uint _getToken1Amount, uint _payToken0Amount, uint _refundToken0Amount) {
        uint _lastStartX = uint(- 1);

        uint _tmpStartX;
        uint _tmpToken0;
        uint _tmpToken1;
        uint _lastIndex;
        for (uint _tmpIndex = lastNodeIndex; _tmpIndex != 0;) {
            if (debtInfos[_tmpIndex].pledgeAmount == 0 && debtInfos[_tmpIndex].repayAmount == 0) {
                _tmpIndex = debtInfos[_tmpIndex].lastIndex;
                continue;
            }
            _tmpStartX = getStartPoint(debtInfos[_tmpIndex].pledgeAmount, debtInfos[_tmpIndex].repayAmount, _reserve0, _reserve1);
            if (_tmpStartX < _reserve1) {
                if (_tmpStartX.add(debtInfos[_tmpIndex].pledgeAmount) <= _reserve1) {
                    _getToken1Amount = _getToken1Amount.add(debtInfos[_tmpIndex].pledgeAmount);
                    _payToken0Amount = _payToken0Amount.add(debtInfos[_tmpIndex].repayAmount);
                    _lastIndex = debtInfos[_tmpIndex].lastIndex;
                    emit DebtUpdate(debtInfos[_tmpIndex].user, debtInfos[_tmpIndex].debtId, 0, 0);
                    if (debtInfos[_tmpIndex].debtToken0Amount == 0 && debtInfos[_tmpIndex].debtToken1Amount == 0) {
                        delDebit(_tmpIndex);
                    } else {
                        totalPledge = totalPledge.sub(debtInfos[_tmpIndex].pledgeAmount);
                        debtInfos[_tmpIndex].pledgeAmount = 0;
                        debtInfos[_tmpIndex].repayAmount = 0;
                    }

                    _tmpIndex = _lastIndex;
                    continue;
                }
                _tmpToken1 = _reserve1.sub(_tmpStartX);
                _tmpToken0 = _reserve0.mul(_tmpToken1)/ _tmpStartX;
                _getToken1Amount = _getToken1Amount.add(_tmpToken1);
                _payToken0Amount = _payToken0Amount.add(_tmpToken0);
                debtInfos[_tmpIndex].pledgeAmount = debtInfos[_tmpIndex].pledgeAmount.sub(_tmpToken1);
                debtInfos[_tmpIndex].repayAmount = debtInfos[_tmpIndex].repayAmount.sub(_tmpToken0);

                totalPledge = totalPledge.sub(_tmpToken1);
                _tmpStartX = _reserve1;
                emit DebtUpdate(debtInfos[_tmpIndex].user, debtInfos[_tmpIndex].debtId, debtInfos[_tmpIndex].pledgeAmount, debtInfos[_tmpIndex].repayAmount);
            }

            if (_tmpStartX >= _lastStartX) {
                _payToken0Amount = _payToken0Amount.add(debtInfos[_tmpIndex].repayAmount);
                _refundToken0Amount = _refundToken0Amount.add(debtInfos[_tmpIndex].repayAmount);
                debtInfos[_tmpIndex].debtToken0Amount = debtInfos[_tmpIndex].debtToken0Amount.add(debtInfos[_tmpIndex].repayAmount);
                debtInfos[_tmpIndex].debtToken1Amount = debtInfos[_tmpIndex].debtToken1Amount.add(debtInfos[_tmpIndex].pledgeAmount);

                totalPledge = totalPledge.sub(debtInfos[_tmpIndex].pledgeAmount);
                debtInfos[_tmpIndex].pledgeAmount = 0;
                debtInfos[_tmpIndex].repayAmount = 0;
                _tmpIndex = debtInfos[_tmpIndex].lastIndex;

                resort(_tmpIndex, _reserve0, _reserve1);
                continue;
            }

            if (_tmpStartX.add(debtInfos[_tmpIndex].pledgeAmount) > _lastStartX) {
                _tmpToken1 = _tmpStartX.add(debtInfos[_tmpIndex].pledgeAmount).sub(_lastStartX);
                // y = r0 * r1 * (tmpx + p - lastx) / (lastx * (tmpx + p))
                _tmpToken0 = _reserve0.mul(_reserve1).mul(_tmpToken1) / (_lastStartX.mul(_tmpStartX.add(debtInfos[_tmpIndex].pledgeAmount)));
                _payToken0Amount = _payToken0Amount.add(_tmpToken0);
                _refundToken0Amount = _refundToken0Amount.add(_tmpToken0);
                debtInfos[_tmpIndex].debtToken0Amount = debtInfos[_tmpIndex].debtToken0Amount.add(_tmpToken0);
                debtInfos[_tmpIndex].debtToken1Amount = debtInfos[_tmpIndex].debtToken1Amount.add(_tmpToken1);
                debtInfos[_tmpIndex].pledgeAmount = debtInfos[_tmpIndex].pledgeAmount.sub(_tmpToken1);
                debtInfos[_tmpIndex].repayAmount = debtInfos[_tmpIndex].repayAmount.sub(_tmpToken0);
                totalPledge = totalPledge.sub(_tmpToken1);
            }
            _lastStartX = _tmpStartX;
            _tmpIndex = debtInfos[_tmpIndex].lastIndex;
        }
    }

    function resort(uint _tmpIndex, uint _reserve0, uint _reserve1) internal {
        uint _x = getStartPoint(debtInfos[_tmpIndex].debtToken1Amount, debtInfos[_tmpIndex].debtToken0Amount, _reserve0, _reserve1);
        uint _sellPoint;
        uint _index = 0;
        for (; debtInfos[_index].nextIndex != 0; _index = debtInfos[_index].nextIndex) {
            if (debtInfos[_index].pledgeAmount == 0 && debtInfos[_index].repayAmount == 0) {
                _sellPoint = getStartPoint(debtInfos[_index].debtToken1Amount, debtInfos[_index].debtToken0Amount, _reserve0, _reserve1);
            } else {
                _sellPoint = getStartPoint(debtInfos[_index].pledgeAmount, debtInfos[_index].repayAmount, _reserve0, _reserve1);
            }

            if (_x <= _sellPoint) {
                debtInfos[_tmpIndex].lastIndex = debtInfos[_index].lastIndex;
                debtInfos[_tmpIndex].nextIndex = _index;
                debtInfos[_index].lastIndex = _tmpIndex;
                debtInfos[debtInfos[_tmpIndex].lastIndex].nextIndex = _tmpIndex;
                return;
            }
        }
        debtInfos[_tmpIndex].lastIndex = _index;
        debtInfos[_index].nextIndex = _tmpIndex;
        debtInfos[_tmpIndex].nextIndex = 0;
        return;
    }

    function getStartPoint(uint _pledgeAmount, uint _targetNum, uint _reserve0, uint _reserve1) internal pure returns (uint){
        uint _k = _reserve0.mul(_reserve1);

        uint _tmpMul = _k * _pledgeAmount;
        uint _res;
        if (_tmpMul / _pledgeAmount != _k) {
            _res = Math.mulDiv(_k, _pledgeAmount, _targetNum);
        } else {
            _res = _tmpMul / _targetNum;
        }
        return (Math.sqrt(_res.add(_pledgeAmount.mul(_pledgeAmount) / 4))).sub(_pledgeAmount / 2);
    }

    function insertDebit(uint _lastIndex, uint _pledgeAmount, uint _targetNum, address _to) internal returns (uint _debtId) {
        uint _insertIndex = getNullIndex();
        _debtId = uniqueDebtId;
        uniqueDebtId++;
        debtInfo memory _info = debtInfo({
        user : _to,
        debtId : _debtId,
        pledgeAmount : _pledgeAmount,
        repayAmount : _targetNum,
        debtToken0Amount : 0,
        debtToken1Amount : 0,
        nextIndex : debtInfos[_lastIndex].nextIndex,
        lastIndex : _lastIndex
        });

        if (_info.nextIndex != 0) {
            debtInfos[_info.nextIndex].lastIndex = _insertIndex;
        } else {
            lastNodeIndex = _insertIndex;
        }
        debtInfos[_lastIndex].nextIndex = _insertIndex;
        if (_insertIndex == debtInfos.length) {
            debtInfos.push(_info);
        } else {
            debtInfos[_insertIndex] = _info;
        }
        totalPledge = totalPledge.add(_pledgeAmount);
    }

    function getNullIndex() private view returns (uint _i) {
        _i = 1;
        for (; _i < debtInfos.length; _i++) {
            if (debtInfos[_i].debtId == 0) {
                return _i;
            }
        }
    }

    function delDebit(uint _index) private {
        if (_index == lastNodeIndex) {
            lastNodeIndex = debtInfos[_index].lastIndex;
        }
        debtInfos[debtInfos[_index].lastIndex].nextIndex = debtInfos[_index].nextIndex;
        if (debtInfos[_index].nextIndex != 0) {
            debtInfos[debtInfos[_index].nextIndex].lastIndex = debtInfos[_index].lastIndex;
        }
        totalPledge = totalPledge.sub(debtInfos[_index].pledgeAmount);
        delete debtInfos[_index];
    }
}
