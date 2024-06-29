// # Write your MySQL query statement below
// SELECT product_id, IFNULL(ROUND( SUM(total) / SUM(units), 2), 0) AS average_price FROM (
//     SELECT p.product_id, price * units AS total, units FROM prices p
//     LEFT JOIN unitssold u ON p.product_id = u.product_id AND
//     p.start_date <= u.purchase_date AND p.end_date >= u.purchase_date
// ) AS products
// GROUP BY products.product_id

