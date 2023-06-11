/**
 * @param {Array} arr1
 * @param {Array} arr2
 * @return {Array}
 * https://leetcode.com/problems/join-two-arrays-by-id/
 */
var join = function (arr1, arr2) {
    const obj = {};
    arr1.forEach(a => obj[a.id] = a);
    arr2.forEach(a => obj[a.id] = { ...obj[a.id], ...a });
    return Object.values(obj);
};