# https://leetcode.com/problems/fix-names-in-a-table/
# Write your MySQL query statement below
SELECT
    user_id,
    CONCAT(
        UPPER(LEFT(NAME, 1)),
        LOWER(RIGHT(NAME, LENGTH(NAME) -1))
    ) AS name
FROM
    Users
ORDER BY
    user_id