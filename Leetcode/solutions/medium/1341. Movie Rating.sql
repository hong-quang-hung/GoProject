# https://leetcode.com/problems/movie-rating/
# Write your MySQL query statement below
# Write your MySQL query statement below
(
    SELECT
        a.name AS results
    FROM
        Users a
        INNER JOIN (
            SELECT
                user_id,
                COUNT(movie_id) AS count
            FROM
                MovieRating
            GROUP BY
                user_id
        ) b ON a.user_id = b.user_id
    ORDER BY
        b.count DESC,
        a.name
    LIMIT
        1
)
UNION
ALL (
    SELECT
        a.title
    FROM
        Movies a
        INNER JOIN (
            SELECT
                movie_id,
                AVG(rating) AS avg
            FROM
                MovieRating
            WHERE
                DATE_FORMAT(created_at, '%Y-%m') = '2020-02'
            GROUP BY
                movie_id
        ) b ON a.movie_id = b.movie_id
    ORDER BY
        b.avg DESC,
        a.title
    LIMIT
        1
)