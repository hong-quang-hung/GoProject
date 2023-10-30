# https://leetcode.com/problems/consecutive-numbers/
# Write your MySQL query statement below
SELECT
    DISTINCT o.num AS ConsecutiveNums
FROM
    LOGS o
    JOIN LOGS tw
    JOIN LOGS th ON o.id + 1 = tw.id
    AND tw.id + 1 = th.id
    AND o.num = tw.num
    AND tw.num = th.num