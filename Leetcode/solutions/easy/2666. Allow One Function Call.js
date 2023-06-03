/**
 * @param {Function} fn
 * @return {Function}
 * https://leetcode.com/problems/allow-one-function-call/
 */
var once = function (fn) {
    let once = true;
    return function (...args) {
        if (once) {
            once = false;
            return fn(...args);
        }
    }
};