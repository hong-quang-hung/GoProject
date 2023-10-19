# https://leetcode.com/problems/the-number-of-employees-which-report-to-each-employee/
# Write your MySQL query statement below
SELECT
    employee_id,
    name,
    COUNT(employee_id) AS reports_count,
    AVG(average_age) AS average_age
FROM
    Employees
GROUP BY
    reports_to
HAVING
    COUNT(employee_id) >= 0