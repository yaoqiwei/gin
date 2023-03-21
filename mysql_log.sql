select employees_id,if(name like 'M%',salary,0) as bonus 
from 
Employees;

select concat(upper(left(name,1)),lower(right(name,len(name)-1))) as name 
from Users 
order by user_id asc;

select Department.name as Department,e1.name as Employee,salary
from employee e1
left join Department on e1.departmentId = department.id 
where
3 > (SELECT
            COUNT(DISTINCT e2.Salary)
        FROM
            Employee e2
        WHERE
            e2.Salary > e1.Salary
                AND e1.DepartmentId = e2.DepartmentId
        )


