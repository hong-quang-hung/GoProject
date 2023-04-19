/* https://leetcode.com/problems/managers-with-at-least-5-direct-reports/description/ */

/* Write your T-SQL query statement below */
;WITH cte AS (
  SELECT managerId, COUNT(1) AS count FROM EMPLOYEE GROUP BY managerId HAVING COUNT(*)>=5
)
SELECT name FROM EMPLOYEE WHERE id IN (SELECT managerid FROM cte)