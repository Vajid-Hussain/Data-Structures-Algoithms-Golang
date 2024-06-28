// -- Write your PostgreSQL query statement below
// SELECT CASE WHEN e.employee_id IS NULL THEN s.employee_id ELSE e.employee_id END AS employee_id FROM employees e
// FULL JOIN salaries s ON e.employee_id = s.employee_id
// WHERE s.employee_id IS NULL OR e.employee_id IS NULL

// ORDER BY 1