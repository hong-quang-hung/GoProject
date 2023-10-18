# https://leetcode.com/problems/rising-temperature/ 
# Write your T-SQL query statement below
SELECT
    b.id
FROM
    Weather a
    JOIN Weather b ON ADDDATE(a.recordDate, 1) = b.recordDate
    AND a.temperature < b.temperature