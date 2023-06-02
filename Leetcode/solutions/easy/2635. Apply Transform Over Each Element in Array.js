/**
 * @param {number[]} arr
 * @param {Function} fn
 * @return {number[]}
 * https://leetcode.com/problems/apply-transform-over-each-element-in-array/
 */
var map = function(arr, fn) {
    for (let i = 0; i < arr.length; i++) {
        arr[i] = fn(arr[i], i)
    }
    return arr
};