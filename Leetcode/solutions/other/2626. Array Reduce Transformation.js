/**
 * @param {number[]} nums
 * @param {Function} fn
 * @param {number} init
 * @return {number}
 * https://leetcode.com/problems/array-reduce-transformation/
 */
var reduce = function(nums, fn, init) {
    if (nums.length === 0) return init;
    let aws = init;
    for(var i = 0; i< nums.length; i++) {
        aws = fn(aws, nums[i])
    }
    return aws
};