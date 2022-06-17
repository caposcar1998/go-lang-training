package employee

import (
	"encoding/json"
	"io"
	"net/http"
)

func QueryEmployee(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id := r.URL.Query().Get("id")
		var employee = RetrieveOneValueById("employee", id)
		var returnValues, _ = json.MarshalIndent(employee, "", "  ")
		w.Write(returnValues)

	case "POST":
		var employee Employee
		err := json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var employeeSaved = CreateEmployee(employee)
		if employeeSaved {
			io.WriteString(w, "True")
		} else {
			io.WriteString(w, "False")
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
		if employeeSaved {
			io.WriteString(w, "True")
		} else {
			io.WriteString(w, "False")
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
		var employees, _ = RetrieveByPositions("employee", position)
		var returnValues, _ = json.MarshalIndent(employees, "", "  ")
		w.Write(returnValues)
	}
}
