package main

import (
	"errors"
	"fmt"
	"main/employee"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/employee", employee.QueryEmployee)
	http.HandleFunc("/employees", employee.QueryEmployees)
	http.HandleFunc("/employees/positions", employee.QueryEmployeesPosition)

	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
