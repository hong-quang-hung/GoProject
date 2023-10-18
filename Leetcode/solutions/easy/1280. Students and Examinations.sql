/* https://leetcode.com/problems/students-and-examinations/ */
# Write your MySQL query statement below
WITH FullJoin AS (
    SELECT
        a.student_id,
        a.student_name,
        b.subject_name
    FROM
        Students a,
        Subjects b
),
Attended AS (
    SELECT
        student_id,
        subject_name,
        COUNT(1) AS count
    FROM
        Examinations
    GROUP BY
        student_id,
        subject_name
)
SELECT
    a.student_id,
    a.student_name,
    a.subject_name,
    IFNULL(b.count, 0) AS attended_exams
FROM
    FullJoin a
    LEFT JOIN Attended b ON a.student_id = b.student_id
    AND a.subject_name = b.subject_name
ORDER BY
    a.student_id,
    a.subject_name