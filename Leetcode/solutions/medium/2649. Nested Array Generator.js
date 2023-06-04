/**
 * @param {Array} arr
 * @return {Generator}
 * https://leetcode.com/problems/nested-array-generator/
 */
var inorderTraversal = function* (arr) {
  yield* arr.flat(Infinity);
};