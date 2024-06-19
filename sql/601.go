// -- -- Write your PostgreSQL query statement below
// -- SELECT id, visit_date, people FROM stadium s1
// -- WHERE people >= 100 AND
// -- ( 2 = (
// --     SELECT COUNT(*) FROM stadium s2
// --     WHERE  people >=100 AND id IN (s1.id+1, s1.id+2) 
// -- )
// -- OR 
// -- 2 = (
// --     SELECT COUNT(*) FROM stadium s3
// --     WHERE people >= 100 AND id IN (s1.id-1, s1.id-2)
// -- ) 
// -- OR
// -- 2 = (
// --     SELECT COUNT(*) FROM stadium s4 
// --     WHERE people >= 100 AND id IN (s1.id+1, s1.id-1)
// -- ) )

// WITH helper AS (
//     SELECT id, visit_date, people, 
//         id - ROW_NUMBER() OVER(PARTITION BY people >=100 ORDER BY id) AS group
//     FROM stadium WHERE people >= 100
// )

// SELECT id, visit_date, people FROM helper h1
// WHERE h1.group IN (
//     SELECT h2.group FROM helper h2 
//     GROUP BY h2.group
//     HAVING COUNT(*) >= 3
// )
