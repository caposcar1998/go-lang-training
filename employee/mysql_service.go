package employee

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

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

func RetrieveOneValueById(table string, value string) Employee {
	var (
		id           int
		full_name    string
		position     int
		salary       float64
		joined       time.Time
		on_probation bool
		created_at   time.Time
	)
	var employee Employee
	db := dbConn()
	row := db.QueryRow("SELECT * FROM " + table + " WHERE id = " + value)
	err := row.Scan(&id, &full_name, &position, &salary, &joined, &on_probation, &created_at)
	if err != nil {
		log.Fatal(err)
	}
	employee = Employee{id, full_name, Position(position), salary, joined, on_probation, created_at}

	return employee
}

func RetrieveByPositions(table string, p string) ([]Employee, error) {
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
	rows, err := db.Query("SELECT * FROM ? WHERE position = ?", table, p)
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
	rows, err := db.Query("SELECT * FROM ? ", table)
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

func CreateEmployee(employee Employee) bool {
	db := dbConn()
	res, err := db.Exec("INSERT INTO employee (full_name, position,salary,joined,on_probation) VALUES (? ,? , ? , ? , ?)", employee.FullName, employee.Position, employee.Salary, time.Now(), employee.OnProbation)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	return true
}

func UpdateEmployee(employee Employee, id string) bool {
	db := dbConn()
	res, err := db.Exec("UPDATE employee SET full_name = ?, position = ?, salary = ?, joined = ?, on_probation = ? WHERE id = ? ", employee.FullName, employee.Position, employee.Salary, time.Now(), employee.OnProbation, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	return true
}
