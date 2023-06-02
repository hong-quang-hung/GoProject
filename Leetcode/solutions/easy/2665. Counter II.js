/**
 * @param {integer} init
 * @return { increment: Function, decrement: Function, reset: Function }
 * https://leetcode.com/problems/counter-ii/
 */
var createCounter = function(init) {
    let n = init;
    let increment = function() {
        return ++n;
    }
    let decrement = function() {
        return --n;
    }
    let reset = function() {
        return n = init;
    }
    return { increment: increment, decrement: decrement, reset: reset }
};

/**
 * const counter = createCounter(5)
 * counter.increment(); // 6
 * counter.reset(); // 5
 * counter.decrement(); // 4
 */