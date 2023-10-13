/* https://leetcode.com/problems/project-employees-i/ */
/* Write your T-SQL query statement below */

SELECT a.project_id, ROUND((SUM(b.experience_years)*1.0)/COUNT(1), 2) AS average_years
FROM Project a
LEFT JOIN Employee b ON a.employee_id = b.employee_id
GROUP BY a.project_id