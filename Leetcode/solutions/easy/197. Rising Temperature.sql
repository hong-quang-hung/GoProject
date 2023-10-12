/* https://leetcode.com/problems/rising-temperature/ */
/* Write your T-SQL query statement below */

SELECT b.id
FROM Weather a
JOIN Weather b ON DATEDIFF(DAY, a.recordDate, b.recordDate) = 1 AND a.temperature < b.temperature