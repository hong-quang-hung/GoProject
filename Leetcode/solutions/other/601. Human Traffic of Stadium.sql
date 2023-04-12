/* https://leetcode.com/problems/human-traffic-of-stadium/ */
/* Write your T-SQL query statement below */

-- create table Stadium (id int, visit_date date, people int)

-- insert into Stadium  select 1, '2017-01-01', 10
-- insert into Stadium  select 2, '2017-01-02', 109
-- insert into Stadium  select 3, '2017-01-03', 150
-- insert into Stadium  select 4, '2017-01-04', 99
-- insert into Stadium  select 5, '2017-01-05', 145
-- insert into Stadium  select 6, '2017-01-06', 1455
-- insert into Stadium  select 7, '2017-01-07', 199
-- insert into Stadium  select 8, '2017-01-09', 1

;WITH get_groups AS
(
	SELECT a.*, a.id - a.rnum AS grp_num
	FROM
	(
		SELECT s.*,
		ROW_NUMBER() OVER(ORDER BY id) AS rnum
		FROM stadium s
		WHERE s.people >= 100
	) a
),
qualifying_groups AS
(
    SELECT grp_num 
	FROM get_groups
	GROUP BY grp_num
	HAVING COUNT(*) >= 3
)
SELECT g.id, g.visit_date, g.people 
FROM get_groups g INNER JOIN qualifying_groups q ON g.grp_num = q.grp_num 
ORDER BY g.visit_date;

-- drop table Stadium 