# https://leetcode.com/problems/game-play-analysis-iv/
# Write your T-SQL query statement below
WITH CTE AS (
    SELECT
        player_id,
        FIRST_VALUE(event_date) OVER (
            PARTITION BY player_id
            ORDER BY
                event_date
        ) AS first_log,
        LEAD(event_date) OVER (
            PARTITION BY player_id
            ORDER BY
                event_date
        ) AS second_log
    FROM
        Activity
)
SELECT
    ROUND(
        COUNT(
            CASE
                WHEN ADDDATE(first_log, 1) = second_log THEN player_id
            END
        ) * 1.0 / COUNT(DISTINCT player_id),
        2
    ) AS fraction
FROM
    CTE