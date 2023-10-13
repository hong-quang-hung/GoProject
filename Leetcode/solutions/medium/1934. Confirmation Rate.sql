/* https://leetcode.com/problems/confirmation-rate/ */
/* Write your T-SQL query statement below */

;WITH Confirmed AS (
    SELECT user_id, CAST(COUNT(user_id) AS NUMERIC(5, 2)) AS count
    FROM Confirmations
    WHERE action = 'confirmed'
    GROUP BY user_id
),
Total AS (
    SELECT user_id, CAST(COUNT(user_id) AS NUMERIC(5, 2)) as count
    FROM Confirmations
    GROUP BY user_id
)
SELECT a.user_id, CASE WHEN ISNULL(c.count, 0) = 0 THEN 0 ELSE ROUND(ISNULL(b.count, 0)/c.count, 2) END AS confirmation_rate
FROM Signups a
LEFT JOIN Confirmed b ON a.user_id = b.user_id
LEFT JOIN Total c ON a.user_id = c.user_id;