/* https://leetcode.com/problems/monthly-transactions-i/ */
/* Write your T-SQL query statement below */

SELECT CONVERT(CHAR(7), trans_date , 126) AS month, country
    , COUNT(1) AS trans_count
    , SUM(CASE WHEN state = 'approved' THEN 1 ELSE 0 END) AS approved_count
    , SUM(amount) AS trans_total_amount
    , SUM(CASE WHEN state = 'approved' THEN amount ELSE 0 END) AS approved_total_amount
FROM Transactions
GROUP BY country, CONVERT(CHAR(7), trans_date, 126)