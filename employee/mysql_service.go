package employee

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Repository interface {
	Employees(ctx context.Context, pos Position) ([]Employee, error)
	EmployeeRetrieve(ctx context.Context, id int) (*Employee, error)
	Save(ctx context.Context, e *Employee) error
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "123456"
	dbName := "Beat"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func (e Employee) EmployeeRetrieve(id int) (Employee, error) {
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
	db := dbConn()
	row := db.QueryRow("SELECT * FROM employee WHERE id = ? ", id)
	err := row.Scan(&idR, &full_name, &position, &salary, &joined, &on_probation, &created_at)
	if err != nil {
		log.Fatal(err)
	}
	employee = Employee{idR, full_name, Position(position), salary, joined, on_probation, created_at}

	return employee, nil
}

func (e Employee) Employees(p string) ([]Employee, error) {
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
	db := dbConn()
	rows, err := db.Query("SELECT * FROM employee WHERE position = ?", p)
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

func RetrieveAllValues(table string) ([]Employee, error) {
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
	db := dbConn()
	rows, err := db.Query("SELECT * FROM " + table)
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

func (e Employee) Save(employee *Employee) error {
	db := dbConn()
	_, err := db.Exec("INSERT INTO employee (full_name, position,salary,joined,on_probation) VALUES (? ,? , ? , ? , ?)", employee.FullName, employee.Position, employee.Salary, time.Now(), employee.OnProbation)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func UpdateEmployee(employee Employee, id string) error {
	db := dbConn()
	_, err := db.Exec("UPDATE employee SET full_name = ?, position = ?, salary = ?, joined = ?, on_probation = ? WHERE id = ? ", employee.FullName, employee.Position, employee.Salary, time.Now(), employee.OnProbation, id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
