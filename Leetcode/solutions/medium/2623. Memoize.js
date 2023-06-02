/**
 * @param {Function} fn
 * https://leetcode.com/problems/memoize/
 */
function memoize(fn) {
    let check = new Map();
    return function(...args) {
        let k = JSON.stringify(args);
        if (!check.has(k)) {
            check.set(k, fn(...args));
        }
        return check.get(k);
    }
}

/** 
 * let callCount = 0;
 * const memoizedFn = memoize(function (a, b) {
 *	 callCount += 1;
 *   return a + b;
 * })
 * memoizedFn(2, 3) // 5
 * memoizedFn(2, 3) // 5
 * console.log(callCount) // 1 
 */