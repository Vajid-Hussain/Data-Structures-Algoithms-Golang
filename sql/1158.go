// -- -- -- Write your PostgreSQL query statement below
// -- SELECT user_id AS buyer_id, join_date,
// -- SUM(CASE WHEN  TO_CHAR(o.order_date, 'YYYY') = '2019' THEN 1 ELSE 0 END ) AS orders_in_2019
// -- FROM users u 
// -- LEFT JOIN orders o ON o.buyer_id = u.user_id 
// -- GROUP BY u.user_id, u.join_date
// -- order by u.user_id


// WITH temp AS (
//     SELECT u.user_id AS buyer_id, u.join_date, COUNT(*) AS orders_in_2019 FROM users u 
//     INNER JOIN orders o ON o.buyer_id  = u.user_id
//     WHERE EXTRACT ('year' FROM o.order_date) = '2019'
//     GROUP BY u.user_id, u.join_date
// )

// SELECT u.user_id buyer_id, u.join_date, COALESCE(orders_in_2019, 0) orders_in_2019 FROM temp
// RIGHT JOIN users u ON u.user_id = temp.buyer_id
