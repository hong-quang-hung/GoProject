
/**
 * @param {Function[]} functions
 * @param {number} n
 * @return {Function}
 * https://leetcode.com/problems/promise-pool/
 */
var promisePool = async function (functions, n) {
    async function evaluateNext() {
        if (functions.length === 0) return;
        const fn = functions.shift();
        await fn();
        await evaluateNext();
    }
    const nPromises = Array(n).fill().map(evaluateNext);
    await Promise.all(nPromises);
};