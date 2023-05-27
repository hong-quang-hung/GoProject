/**
 * @param {any[]} arr
 * @param {number} depth
 * @return {any[]}
 * https://leetcode.com/problems/flatten-deeply-nested-array/
 */
var flat = function (arr, n) {
    let res = [];
    const recursion = (nums, l) => {
      for (const num of nums) {
        if (Array.isArray(num) && l > 0) {
          recursion(num, l - 1);
        } else {
          res.push(num);
        }
      }
    }

    recursion(arr, n);
    return res;
};