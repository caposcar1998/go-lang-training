package employee

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func QueryEmployee(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var rep Repository
		ctx := context.TODO()
		id := r.URL.Query().Get("id")

		var employee, err = rep.EmployeeRetrieve(ctx, id)
		if err != nil {
			var returnValues, _ = json.MarshalIndent(employee, "", "  ")
			w.Write(returnValues)
		} else {
			io.WriteString(w, err.Error())
		}

	case "POST":
		var rep Repository
		ctx := context.TODO()
		var employee Employee
		err := json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var employeeSaved = rep.Save(ctx, &employee)
		if employeeSaved != nil {
			io.WriteString(w, employeeSaved.Error())
		} else {
			io.WriteString(w, "ok")
		}

	case "PUT":
		var employee Employee
		id := r.URL.Query().Get("id")
		err := json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var employeeSaved = UpdateEmployee(employee, id)
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
		var employees, _ = RetrieveAllValues("employee")
		var returnValues, _ = json.MarshalIndent(employees, "", "  ")
		w.Write(returnValues)
	}
}

func QueryEmployeesPosition(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		position := r.URL.Query().Get("position")
		var employees, _ = Employees(position)
		var returnValues, _ = json.MarshalIndent(employees, "", "  ")
		w.Write(returnValues)
	}
}
