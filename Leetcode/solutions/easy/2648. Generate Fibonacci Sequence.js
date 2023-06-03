/**
 * @return {Generator<number>}
 * https://leetcode.com/problems/generate-fibonacci-sequence/
 */
var fibGenerator = function*() {
    let fib = [0, 1];
    yield fib[0];
    yield fib[1];
    while (true) {
        let next = fib[0] + fib[1];
        yield next;
        fib[0] = fib[1];
        fib[1] = next;
    }
};