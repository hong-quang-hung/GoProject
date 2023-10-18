# https://leetcode.com/problems/human-traffic-of-stadium/
# Write your MySQL query statement below
WITH get_groups AS (
	SELECT
		a.*,
		a.id - a.rnum AS grp_num
	FROM
		(
			SELECT
				s.*,
				ROW_NUMBER() OVER(
					ORDER BY
						id
				) AS rnum
			FROM
				stadium s
			WHERE
				s.people >= 100
		) a
),
qualifying_groups AS (
	SELECT
		grp_num
	FROM
		get_groups
	GROUP BY
		grp_num
	HAVING
		COUNT(*) >= 3
)
SELECT
	g.id,
	g.visit_date,
	g.people
FROM
	get_groups g
	INNER JOIN qualifying_groups q ON g.grp_num = q.grp_num
ORDER BY
	g.visit_date;