/**
 * @param {any} object
 * @return {string}
 * https://leetcode.com/problems/convert-object-to-json-string/
 */
var jsonStringify = function (object) {
    return typeof object === 'string' ? '"' + object + '"' :
        object === null || typeof object !== 'object' ? object :
            Array.isArray(object) ? '[' + object.reduce((acc, x) => acc + jsonStringify(x) + ',', '').slice(0, -1) + ']' :
                '{' + Object.entries(object).reduce((acc, x) => acc + jsonStringify(x[0]) + ':' + jsonStringify(x[1]) + ',', '').slice(0, -1) + '}';
};