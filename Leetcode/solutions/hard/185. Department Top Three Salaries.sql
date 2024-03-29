# https://leetcode.com/problems/department-top-three-salaries/
# Write your MySQL query statement below
SELECT
    a.name AS Department,
    b.name AS Employee,
    b.salary AS Salary
FROM
    Department a
    INNER JOIN (
        SELECT
            *,
            DENSE_RANK() OVER(
                PARTITION BY departmentId
                ORDER BY
                    salary DESC
            ) AS salary_rank
        FROM
            Employee
    ) b ON a.id = b.departmentId
    AND b.salary_rank <= 3