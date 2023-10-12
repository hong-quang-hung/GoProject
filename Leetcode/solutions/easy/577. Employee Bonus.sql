/* https://leetcode.com/problems/employee-bonus/ */
/* Write your T-SQL query statement below */

SELECT a.name, b.bonus
FROM Employee a
LEFT JOIN Bonus b ON a.empId = b.empId
WHERE ISNULL(b.bonus, 0) < 1000