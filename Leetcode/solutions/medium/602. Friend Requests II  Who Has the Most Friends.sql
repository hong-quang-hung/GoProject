# https://leetcode.com/problems/friend-requests-ii-who-has-the-most-friends/
# Write your MySQL query statement below
SELECT
    id,
    SUM(id)
FROM
    (
        (
            SELECT
                requester_id AS id,
                COUNT(requester_id) AS num
            FROM
                RequestAccepted
            GROUP BY
                requester_id
        )
        UNION
        ALL (
            SELECT
                accepter_id AS id,
                COUNT(accepter_id) AS num
            FROM
                RequestAccepted
            GROUP BY
                accepter_id
        )
    ) a
GROUP BY
    id
ORDER BY
    num DESC
LIMIT
    1