package employee

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type repository struct {
	db *sql.DB
}

func (r *repository) Close() {
	r.db.Close()
}

func (r *repository) EmployeeRetrieve(id int) (Employee, error) {
	var (
		idR          int
		full_name    string
		position     int
		salary       float64
		joined       time.Time
		on_probation bool
		created_at   time.Time
	)

	var employee Employee
	row := r.db.QueryRow("SELECT * FROM employee WHERE id = ? ", id)
	err := row.Scan(&idR, &full_name, &position, &salary, &joined, &on_probation, &created_at)
	if err != nil {
		log.Fatal(err)
	}
	employee = Employee{idR, full_name, Position(position), salary, joined, on_probation, created_at}

	return employee, nil
}

func (r *repository) Employees(p string) ([]Employee, error) {
	var (
		id           int
		full_name    string
		position     Position
		salary       float64
		joined       time.Time
		on_probation bool
		created_at   time.Time
	)
	var employees []Employee
	rows, err := r.db.Query("SELECT * FROM employee WHERE position = ?", p)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &full_name, &position, &salary, &joined, &on_probation, &created_at)
		if err != nil {
			log.Fatal(err)
		}
		employees = append(employees, Employee{id, full_name, Position(position), salary, joined, on_probation, created_at})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return employees, nil
}

func (r *repository) RetrieveAllValues(table string) ([]Employee, error) {
	var (
		id           int
		full_name    string
		position     int
		salary       float64
		joined       time.Time
		on_probation bool
		created_at   time.Time
	)
	var employees []Employee
	rows, err := r.db.Query("SELECT * FROM " + table)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &full_name, &position, &salary, &joined, &on_probation, &created_at)
		if err != nil {
			log.Fatal(err)
		}
		employees = append(employees, Employee{id, full_name, Position(position), salary, joined, on_probation, created_at})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return employees, nil
}

func (r *repository) Save(employee *Employee) error {

	_, err := r.db.Exec("INSERT INTO employee (full_name, position,salary,joined,on_probation) VALUES (? ,? , ? , ? , ?)", employee.FullName, employee.Position, employee.Salary, time.Now(), employee.OnProbation)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (r *repository) UpdateEmployee(employee Employee, id string) error {

	_, err := r.db.Exec("UPDATE employee SET full_name = ?, position = ?, salary = ?, joined = ?, on_probation = ? WHERE id = ? ", employee.FullName, employee.Position, employee.Salary, time.Now(), employee.OnProbation, id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
