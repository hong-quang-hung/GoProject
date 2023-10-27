# https://leetcode.com/problems/the-number-of-employees-which-report-to-each-employee/
# Write your MySQL query statement below
SELECT
    a.employee_id,
    a.name,
    b.reports_count,
    b.average_age
FROM
    Employees a
    INNER JOIN (
        SELECT
            reports_to AS employee_id,
            COUNT(employee_id) AS reports_count,
            ROUND(AVG(age), 0) AS average_age
        FROM
            Employees
        WHERE
            reports_to IS NOT NULL
        GROUP BY
            reports_to
        HAVING
            COUNT(employee_id) >= 0
    ) b ON a.employee_id = b.employee_id
ORDER BY
    a.employee_id