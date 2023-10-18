# https://leetcode.com/problems/biggest-single-number/
# Write your MySQL query statement below
SELECT
    a.num AS num
FROM
    (
        SELECT
            num,
            COUNT(num) AS appearance
        FROM
            MyNumbers
        GROUP BY
            num
    ) AS a
WHERE
    a.appearance = 1