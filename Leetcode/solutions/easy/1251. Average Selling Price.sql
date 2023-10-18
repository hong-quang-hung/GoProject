# https://leetcode.com/problems/average-selling-price/
# Write your MySQL query statement below
WITH Total AS (
    SELECT
        a.product_id,
        (a.price * b.units * 1.0) AS prices,
        b.units
    FROM
        Prices a
        LEFT JOIN UnitsSold b ON a.product_id = b.product_id
        AND b.purchase_date >= a.start_date
        AND b.purchase_date <= a.end_date
)
SELECT
    product_id,
    IFNULL(ROUND(SUM(prices) / SUM(units), 2), 0) AS average_price
FROM
    Total
GROUP BY
    product_id