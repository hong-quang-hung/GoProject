# https://leetcode.com/problems/the-number-of-employees-which-report-to-each-employee/
# Write your MySQL query statement below
SELECT
    a.employee_id,
    a.name,
    COUNT(b.employee_id) AS reports_count,
    round(avg(b.age)) AS average_age
FROM
    Employees a
    INNER JOIN Employees b ON a.employee_id = b.reports_to
GROUP BY
    a.employee_id
ORDER BY
    a.employee_id