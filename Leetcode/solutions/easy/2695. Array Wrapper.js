/**
 * @param {number[]} nums
 * https://leetcode.com/problems/array-wrapper/
 */
var ArrayWrapper = function(nums) {
    this.nums = nums;
};

ArrayWrapper.prototype.valueOf = function() {
    return this.nums.reduce((a, b) => a + b, 0);
}

ArrayWrapper.prototype.toString = function() {
    return `[${this.nums}]`;
}