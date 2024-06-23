// -- Write your PostgreSQL query statement below
// SELECT sell_date, COUNT(DISTINCT(product)) AS num_sold, array_to_string(array_agg(DISTINCT(product) ORDER BY product) ,',') AS products
// FROM activities
// GROUP BY sell_date
// ORDER BY sell_date;
