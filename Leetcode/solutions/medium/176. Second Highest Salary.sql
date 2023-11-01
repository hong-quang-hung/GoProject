# https://leetcode.com/problems/second-highest-salary/
# Write your MySQL query statement below
SELECT
    (
        SELECT
            DISTINCT Salary
        FROM
            Employee
        ORDER BY
            salary DESC
        LIMIT
            1 OFFSET 1
    ) AS SecondHighestSalary;