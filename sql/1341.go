// SELECT name AS results FROM (
//     SELECT name, COUNT(*) AS count FROM users
//     INNER JOIN movierating USING (user_id) 
//     GROUP BY user_id, name
//     ORDER BY count DESC, name  
//     LIMIT 1
// )

// UNION ALL

// SELECT title FROM (
//     SELECT title, AVG(rating)  FROM movierating 
//     INNER JOIN movies USING(movie_id) 
//     WHERE to_char(created_at, 'YYYY-MM') = '2020-02'
//     GROUP BY movie_id, title
//     ORDER BY 2 DESC, title
//     LIMIT 1
// )
