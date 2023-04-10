CREATE TABLE Insurance(pid INT, tiv_2015 FLOAT, tiv_2016 FLOAT, lat FLOAT, lon FLOAT)

INSERT INTO Insurance (pid, tiv_2015, tiv_2016, lat, lon) VALUES ('1', '10', '5', '10', '10')
INSERT INTO Insurance (pid, tiv_2015, tiv_2016, lat, lon) VALUES ('2', '20', '20', '20', '20')
INSERT INTO Insurance (pid, tiv_2015, tiv_2016, lat, lon) VALUES ('3', '10', '30', '20', '20')
INSERT INTO Insurance (pid, tiv_2015, tiv_2016, lat, lon) VALUES ('4', '10', '40', '40', '40')

/* Code */

;WITH 
located AS 
(
	SELECT lat AS lat, lon AS lon FROM Insurance GROUP BY lat, lon HAVING COUNT(*) = 1
),
sameValue AS
(
	SELECT tiv_2015 AS value FROM Insurance GROUP BY tiv_2015 HAVING COUNT(*) > 1
)
SELECT ROUND(SUM(tiv_2016), 2) AS tiv_2016 FROM Insurance a JOIN located b ON a.lat = b.lat AND a.lon = b.lon JOIN sameValue c ON a.tiv_2015 = c.value

/* Code */

DROP TABLE Insurance