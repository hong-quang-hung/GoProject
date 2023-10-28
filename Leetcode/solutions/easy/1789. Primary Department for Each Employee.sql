# https://leetcode.com/problems/primary-department-for-each-employee/
# Write your MySQL query statement below
SELECT
    employee_id,
    department_id
FROM
    (
        SELECT
            employee_id,
            department_id,
            primary_flag,
            COUNT(employee_id) OVER(PARTITION BY employee_id) AS Count
        FROM
            Employee
    ) a
WHERE
    Count = 1
    OR primary_flag = 'Y'