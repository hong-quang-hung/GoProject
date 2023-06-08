/**
 * @param {Function} fn
 * @param {Array} args
 * @param {number} t
 * @return {Function}
 * https://leetcode.com/problems/interval-cancellation/
 */
var cancellable = function (fn, args, t) {
    fn(...args);
    let set = setInterval(() => { fn(...args); }, t);
    return () => { clearInterval(set); }
};