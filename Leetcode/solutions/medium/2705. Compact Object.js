/**
 * @param {Object} obj
 * @return {Object}
 * https://leetcode.com/problems/compact-object/
 */
var compactObject = function (obj) {
    if (Array.isArray(obj)) {
        return obj.filter(Boolean).map(compactObject);
    }

    if (obj === null || typeof obj !== 'object') {
        return obj;
    }

    let res = {};
    for (const key in obj) {
        const value = compactObject(obj[key]);
        if (Boolean(value)) {
            res[key] = value;
        }
    }
    return res;
};