/* https://leetcode.com/problems/replace-employee-id-with-the-unique-identifier/ */
/* Write your T-SQL query statement below */

SELECT b.unique_id, a.name
FROM Employees a LEFT JOIN EmployeeUNI b ON a.id = b.id