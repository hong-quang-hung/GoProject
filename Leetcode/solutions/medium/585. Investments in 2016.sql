# https://leetcode.com/problems/investments-in-2016/
# Write your MySQL query statement below
WITH located AS (
	SELECT
		lat AS lat,
		lon AS lon
	FROM
		Insurance
	GROUP BY
		lat,
		lon
	HAVING
		COUNT(*) = 1
),
sameValue AS (
	SELECT
		tiv_2015 AS value
	FROM
		Insurance
	GROUP BY
		tiv_2015
	HAVING
		COUNT(*) > 1
)
SELECT
	ROUND(SUM(tiv_2016), 2) AS tiv_2016
FROM
	Insurance a
	INNER JOIN located b ON a.lat = b.lat
	AND a.lon = b.lon
	INNER JOIN sameValue c ON a.tiv_2015 = c.value