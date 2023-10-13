/* https://leetcode.com/problems/percentage-of-users-attended-a-contest */
/* Write your T-SQL query statement below */

SELECT a.contest_id, ROUND(COUNT(1)*100.0/b.total, 2) AS percentage
FROM Register a, (SELECT COUNT(1) AS total FROM Users) b
GROUP BY a.contest_id
ORDER BY percentage DESC, a.contest_id