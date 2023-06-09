class Calculator {
    /** 
     * @param {number} value
     * https://leetcode.com/problems/calculator-with-method-chaining/
     */
    constructor(value) {
        this.res = value;
    }

    /** 
     * @param {number} value
     * @return {Calculator}
     */
    add(value) {
        this.res += value;
        return this.res;
    }

    /** 
     * @param {number} value
     * @return {Calculator}
     */
    subtract(value) {
        this.res -= value;
        return this.res;
    }

    /** 
     * @param {number} value
     * @return {Calculator}
     */
    multiply(value) {
        this.res *= value;
        return this.res;
    }

    /** 
     * @param {number} value
     * @return {Calculator}
     */
    divide(value) {
        if (value == 0) this.res = "Division by zero is not allowed";
        else this.res /= value;
        return this.res;
    }

    /** 
     * @param {number} value
     * @return {Calculator}
     */
    power(value) {
        this.res = Math.pow(this.res, value);
        return this.res;
    }

    /** 
     * @return {number}
     */
    getResult() {
        return this.res;
    }
}