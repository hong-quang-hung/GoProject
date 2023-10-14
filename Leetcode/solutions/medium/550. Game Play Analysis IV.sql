/* https://leetcode.com/problems/game-play-analysis-iv/ */
/* Write your T-SQL query statement below */

;WITH CTE AS (
    SELECT player_id
        , ROW_NUMBER() OVER (PARTITION BY player_id, device_id ORDER BY event_date) AS rank
    FROM Activity
)
SELECT ROUND(COUNT(1)*1.0/(SELECT COUNT(DISTINCT player_id) FROM Activity), 2) AS fraction
FROM CTE
WHERE rank = 2