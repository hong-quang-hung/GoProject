# https://leetcode.com/problems/find-followers-count/
# Write your MySQL query statement below
SELECT
    a.user_id,
    a.followers_count
FROM
    (
        SELECT
            user_id,
            COUNT(user_id) AS followers_count
        FROM
            Followers
        GROUP BY
            user_id
    ) AS a
ORDER BY
    a.user_id