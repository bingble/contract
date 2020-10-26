pragma solidity >=0.5.16;

// a libraries for performing overflow-safe math, courtesy of DappHub (https://github.com/dapphub/ds-math)

library Math {
    function add(uint x, uint y) internal pure returns (uint z) {
        require((z = x + y) >= x, 'ds-math-add-overflow');
    }

    function sub(uint x, uint y) internal pure returns (uint z) {
        require((z = x - y) <= x, 'ds-math-sub-underflow');
    }

    function mul(uint x, uint y) internal pure returns (uint z) {
        require(y == 0 || (z = x * y) / y == x, 'ds-math-mul-overflow');
    }

    function min(uint x, uint y) internal pure returns (uint z) {
        z = x < y ? x : y;
    }

    // babylonian method (https://en.wikipedia.org/wiki/Methods_of_computing_square_roots#Babylonian_method)
    function sqrt(uint y) internal pure returns (uint z) {
        if (y > 3) {
            z = y;
            uint x = y / 2 + 1;
            while (x < z) {
                z = x;
                x = (y / x + x) / 2;
            }
        } else if (y != 0) {
            z = 1;
        }
    }

    function fullMul(uint x, uint y) internal pure returns (uint l, uint h) {
        uint mm = mulmod(x, y, uint(- 1));
        l = x * y;
        h = mm - l;
        if (mm < l) h -= 1;
    }

    function mulDiv(uint x, uint y, uint z) internal pure returns (uint) {
        (uint l, uint h) = fullMul(x, y);
        require(h < z);

        uint mm = mulmod(x, y, z);
        if (mm > l) h -= 1;
        l -= mm;

        uint pow2 = z & - z;
        z /= pow2;
        l /= pow2;
        l += h * ((- pow2) / pow2 + 1);

        uint r = 1;
        r *= 2 - z * r;
        r *= 2 - z * r;
        r *= 2 - z * r;
        r *= 2 - z * r;
        r *= 2 - z * r;
        r *= 2 - z * r;
        r *= 2 - z * r;
        r *= 2 - z * r;
        return l * r;
    }
}
