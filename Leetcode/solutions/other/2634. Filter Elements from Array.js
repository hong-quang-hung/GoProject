/**
 * @param {number[]} arr
 * @param {Function} fn
 * @return {number[]}
 * https://leetcode.com/problems/filter-elements-from-array/
 */
var filter = function(arr, fn) {
    let res = [];
    for(let i = 0; i< arr.length; i++ ) {
        if (fn(arr[i], i)) res.push(arr[i])
    }
    return res
};