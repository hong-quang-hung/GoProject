# https://leetcode.com/problems/triangle-judgement/
# Write your MySQL query statement below
SELECT
    x,
    y,
    z,
    IF(
        (x + y <= z)
        OR (y + z <= x)
        OR (x + z <= y),
        'No',
        'YES'
    ) AS triangle
FROM
    Triangle