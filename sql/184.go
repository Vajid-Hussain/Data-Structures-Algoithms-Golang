// -- Write your PostgreSQL query statement below
// SELECT d.name AS Department, e.name Employee, e.salary Salary FROM employee e
// INNER JOIN department d ON e.departmentId = d.id
// WHERE salary = (SELECT MAX(salary) FROM employee e1 WHERE e1.departmentId = e.departmentId)
