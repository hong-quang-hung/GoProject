# https://leetcode.com/problems/immediate-food-delivery-ii/
# Write your MySQL query statement below
WITH CTE AS (
    SELECT
        customer_id,
        CASE
            WHEN order_date = customer_pref_delivery_date THEN 1
            ELSE 0
        END AS immediate,
        ROW_NUMBER() OVER (
            PARTITION BY customer_id
            ORDER BY
                order_date
        ) AS delivery_rank
    FROM
        Delivery
)
SELECT
    ROUND(SUM(immediate) * 1.0 / COUNT(1), 4) * 100 AS immediate_percentage
FROM
    CTE
WHERE
    delivery_rank = 1