/**
 * @param {Object | Array} obj
 * @return {boolean}
 * https://leetcode.com/problems/is-object-empty/
 */
var isEmpty = function(obj) {
    return obj && ((Object.keys(obj).length === 0) || (Array.isArray(obj) && obj.length == 0));
};