/**
 * @param {Function} fn
 * @param {number} t milliseconds
 * @return {Function}
 * https://leetcode.com/problems/debounce/
 */
var debounce = function(fn, t) {
    let timeout;
    return function(...args) {
      clearTimeout(timeout);
      timeout = setTimeout(() => {
        fn(...args)
      }, t);
    };
  };