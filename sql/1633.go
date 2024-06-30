// # Write your MySQL query statement below
// SELECT contest_id, ROUND( COUNT(*)/(
//     SELECT COUNT(*) FROM users
// ) * 100, 2) percentage FROM register 
// GROUP BY contest_id
// ORDER BY percentage DESC, contest_id

// SELECT contest_id ,ROUND( (count(*) / (SELECT count(*) FROM users)::decimal)*100, 2) AS percentage FROM register 
// GROUP BY contest_id 
// ORDER BY  percentage desc, contest_id;