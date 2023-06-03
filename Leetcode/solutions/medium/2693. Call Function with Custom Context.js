/**
 * @param {Object} context
 * @param {any[]} args
 * @return {any}
 * https://leetcode.com/problems/call-function-with-custom-context/
 */
Function.prototype.callPolyfill = function(context, ...args) {
    return this.apply(context, args);
}