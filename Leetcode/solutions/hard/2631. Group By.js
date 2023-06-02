/**
 * @param {Function} fn
 * @return {Array}
 * https://leetcode.com/problems/group-by/
 */
Array.prototype.groupBy = function(fn) {
    return this.reduce((accum, item) => {
      const key = fn(item);
      accum[key] ||= [];
      accum[key].push(item);
      return accum;
    }, {});
  };