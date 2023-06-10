/**
 * @param {string} val
 * @return {Object}
 * https://leetcode.com/problems/to-be-or-not-to-be/
 */
var expect = function (val) {
    var toBe = function (x) {
        if (x === val) return true;
        else throw new Error('Not Equal');
    };

    var notToBe = function (x) {
        if (x !== val) return true;
        else throw new Error('Equal');
    };

    return { toBe, notToBe };
};