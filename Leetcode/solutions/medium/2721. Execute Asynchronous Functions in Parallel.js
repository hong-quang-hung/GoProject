/**
 * @param {Array<Function>} functions
 * @return {Promise<any>}
 * https://leetcode.com/problems/execute-asynchronous-functions-in-parallel/
 */
var promiseAll = function (functions) {
    return new Promise((resolve, reject) => {
        let ans = [], j = functions.length;
        functions.forEach((func, i) => {
            func()
                .then((res) => { (ans[i] = res, --j === 0 && resolve(ans)) })
                .catch(reject)
        })
    })
};