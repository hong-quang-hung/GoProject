/* https://leetcode.com/problems/find-users-with-valid-e-mails/ */
/* Write your T-SQL query statement below */

SELECT user_id, name, mail
FROM Users
WHERE mail LIKE '[a-zA-Z]%@leetcode.com' AND LEFT(mail, LEN(mail) - 13) NOT LIKE '%[^a-zA-Z_0-9_.-]%'