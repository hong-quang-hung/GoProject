/* https://leetcode.com/problems/employees-with-missing-information/ */

/* Write your T-SQL query statement below */
SELECT employee_id FROM Employees WHERE employee_id NOT IN (SELECT employee_id FROM Salaries)
UNION ALL
SELECT employee_id FROM Salaries WHERE employee_id NOT IN (SELECT employee_id FROM Employees)
ORDER BY employee_id ASC 