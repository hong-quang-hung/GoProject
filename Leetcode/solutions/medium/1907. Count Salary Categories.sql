# https://leetcode.com/problems/count-salary-categories/
# Write your MySQL query statement below
SELECT
    'Low Salary' AS category,
    COUNT(account_id) AS accounts_count
FROM
    Accounts
WHERE
    income < 20000
UNION
ALL
SELECT
    'Average Salary',
    COUNT(account_id)
FROM
    Accounts
WHERE
    income >= 20000
    AND income <= 50000
UNION
ALL
SELECT
    'High Salary',
    COUNT(account_id)
FROM
    Accounts
WHERE
    income > 50000