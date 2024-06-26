// WITH blocks AS (
//     SELECT users_id FROM users 
//     WHERE banned= 'Yes'
// ) 

// SELECT request_at AS Day, ROUND(
//     (
//         SELECT COUNT(*) FROM trips t1
//         WHERE t1.request_at = t.request_at AND t1.client_id NOT IN ( SELECT users_id FROM blocks) AND t1.driver_id NOT IN ( SELECT users_id FROM blocks)
//         AND (status = 'cancelled_by_driver' OR status ='cancelled_by_client')
//     )::decimal
// / COUNT(*), 2) AS "Cancellation Rate" FROM trips t
// WHERE client_id NOT IN ( SELECT users_id FROM blocks) AND driver_id NOT IN ( SELECT users_id FROM blocks)
// AND request_at >= '2013-10-01' AND request_at <= '2013-10-03'
// GROUP BY request_at
