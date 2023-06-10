/**
 * @param {Function} fn
 * @param {Array} args
 * @param {number} t
 * @return {Function}
 * https://leetcode.com/problems/execute-cancellable-function-with-delay/
 */
var cancellable = function (fn, args, t) {
    let timeoutHandle = setTimeout(() => fn(...args), t)
    return () => clearTimeout(timeoutHandle)
};