# https://leetcode.com/problems/exchange-seats/
# Write your MySQL query statement below
SELECT
    a.id,
    IFNULL(b.student, a.student) AS student
FROM
    Seat a
    LEFT JOIN (
        SELECT
            CASE
                WHEN id % 2 = 1 THEN id + 1
                ELSE id - 1
            END AS id,
            student
        FROM
            Seat
    ) b ON a.id = b.id
ORDER BY
    a.id