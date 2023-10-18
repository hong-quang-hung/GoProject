# https://leetcode.com/problems/average-time-of-process-per-machine/
# Write your MySQL query statement below
SELECT
    a.machine_id,
    ROUND((b.time - a.time) / a.count, 3) AS processing_time
FROM
    (
        SELECT
            machine_id,
            SUM(timestamp) AS time,
            COUNT(machine_id) AS count
        FROM
            Activity
        WHERE
            activity_type = 'start'
        GROUP BY
            machine_id,
            activity_type
    ) a
    JOIN (
        SELECT
            machine_id,
            SUM(timestamp) AS time
        FROM
            Activity
        WHERE
            activity_type = 'end'
        GROUP BY
            machine_id,
            activity_type
    ) b ON a.machine_id = b.machine_id