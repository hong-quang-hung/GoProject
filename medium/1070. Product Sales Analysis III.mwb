-- https://leetcode.com/problems/product-sales-analysis-iii/
-- Write your MySQL query statement below

WITH First AS (
    SELECT product_id, MIN(year) AS year
    FROM Sales
    GROUP BY product_id
)
SELECT a.product_id, b.year AS first_year, a.quantity, a.price
FROM Sales a
JOIN First b ON a.product_id = b.product_id AND a.year = b.year