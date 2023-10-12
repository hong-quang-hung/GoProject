/* https://leetcode.com/problems/customer-who-visited-but-did-not-make-any-transactions/ */
/* Write your T-SQL query statement below */

SELECT a.customer_id, COUNT(a.customer_id) AS count_no_trans
FROM Visits a
LEFT JOIN Transactions b ON a.visit_id = b.visit_id
WHERE b.transaction_id IS NULL
GROUP BY a.customer_id;