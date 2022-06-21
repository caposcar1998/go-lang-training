package employee

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

var ctx = context.TODO()
var employee Employee

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

func QueryEmployee(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		id := r.URL.Query().Get("id")
		repo := &repository{dbConn()}

		var idInteger, _ = strconv.Atoi(id)
		var employee, _ = repo.EmployeeRetrieve(idInteger)

		var returnValues, _ = json.MarshalIndent(employee, "", "  ")
		w.Write(returnValues)

	case "POST":
		err := json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		repo := &repository{dbConn()}
		var employeeSaved = repo.Save(&employee)
		if employeeSaved != nil {
			io.WriteString(w, employeeSaved.Error())
		} else {
			io.WriteString(w, "ok")
		}

	case "PUT":
		id := r.URL.Query().Get("id")
		err := json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		repo := &repository{dbConn()}
		var employeeSaved = repo.UpdateEmployee(employee, id)
		if employeeSaved != nil {
			io.WriteString(w, employeeSaved.Error())
		} else {
			io.WriteString(w, "ok")
		}
	case "DELETE":
		//NOT IMPLEMENTED
	}
}

func QueryEmployees(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		repo := &repository{dbConn()}
		var employees, _ = repo.RetrieveAllValues("employee")
		var returnValues, _ = json.MarshalIndent(employees, "", "  ")
		w.Write(returnValues)
	}
}

func QueryEmployeesPosition(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		position := r.URL.Query().Get("position")
		repo := &repository{dbConn()}
		var employees, _ = repo.Employees(position)
		var returnValues, _ = json.MarshalIndent(employees, "", "  ")
		w.Write(returnValues)
	}
}
